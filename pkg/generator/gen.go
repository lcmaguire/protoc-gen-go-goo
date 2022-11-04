package generator

import (
	"fmt"
	"strings"

	"github.com/lcmaguire/protoc-gen-go-goo/pkg/connectgo"
	"github.com/lcmaguire/protoc-gen-go-goo/pkg/templates"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Generator holds config for determining what code is generated
type Generator struct {
	Tests     bool   // will generate tests for all RPC methods generated default is true.
	ConnectGo bool   // if true will generate connectGo compatable code default is false.
	Server    bool   // if true will generate a server for the services and methods generated default is false.
	GoModPath string // the path of your generated code, Required for server to correctly import newly generated code.

	// firebase trial
	Firebase bool

	// this is temporary
	ServerTemplate         string
	RegisterServerTemplate string
}

// Run will generate RPC methods for your files.
func (g *Generator) Run(gen *protogen.Plugin) error {
	servicesData := []serviceT{}
	fileInfoMap := make(map[string]FileInfo)
	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}

		for _, v := range f.Services {
			g.generateFilesForService(gen, v, f)
			servicesData = append(servicesData, serviceT{ServiceName: string(v.Desc.Name()), FullName: string(v.Desc.FullName())})
			// todo see if v.GoName is equal
		}

		// todo this may be useful for generateFilesForService.
		// may also be useful to document what the key in this map is.
		fileInfoMap = collectFileData(f, fileInfoMap)
	}

	if g.Server {
		// todo handle this more elegantly.
		g.ServerTemplate = templates.ServerTemplate
		g.RegisterServerTemplate = templates.RegisterServiceTemplate
		if g.ConnectGo {
			g.ServerTemplate = connectgo.ServerTemplate
			g.RegisterServerTemplate = connectgo.ServiceHandleTemplate
		}

		for _, v := range fileInfoMap {
			g.generateServer(gen, v, servicesData)
		}
	}

	return nil
}

func (g *Generator) generateFilesForService(gen *protogen.Plugin, service *protogen.Service, file *protogen.File) (outfiles []*protogen.GeneratedFile) {
	serviceFile := g.generateServiceFile(gen, service, file)
	outfiles = append(outfiles, serviceFile)

	// will create a method for all services
	for _, v := range service.Methods {
		requestType := getParamPKG(v.Input.GoIdent.GoImportPath.String()) + "." + v.Input.GoIdent.GoName
		responseType := getParamPKG(v.Output.GoIdent.GoImportPath.String()) + "." + v.Output.GoIdent.GoName

		// get all unique imports. NOTE this will not be ordered as it is a map so file changes may appear due to imports being in different orders.
		mapImports := map[string]any{string(v.Input.GoIdent.GoImportPath): nil, string(v.Output.GoIdent.GoImportPath): nil}
		imports := ""
		for k := range mapImports {
			imports += fmt.Sprintf("\"%s\"\n", k)
		}

		mData := methodData{
			MethodCaller: genMethodCaller(service.GoName),
			ServiceName:  service.GoName,
			template:     g.getMethodTemplate(v.Desc),
			testTemplate: g.getTestMethodTemplate(v.Desc),
			// add in import paths. (done nicely)
			// add in Pkg name done nicely
			MethodName:   v.GoName,
			FullName:     string(v.Desc.FullName()),
			RequestType:  requestType,
			ResponseType: responseType,
			Imports:      imports,
			methodDesc:   v.Desc,
			Pkg:          getParamPKG(file.GoDescriptorIdent.GoImportPath.String()),
			GoPkgName:    strings.ToLower(service.GoName),
		}
		f := g.genRpcMethod(gen, mData)
		outfiles = append(outfiles, f)
		if g.Tests {

			// todo handle this earlier
			if v.Desc.IsStreamingClient() || v.Desc.IsStreamingServer() {
				// imports connect-go gRPC code. Could not include the connect part here but rather in the template.
				// probably should be a field in methodData gRPC import.
				mData.Imports += protogen.GoIdent{GoImportPath: file.GoDescriptorIdent.GoImportPath + "connect"}.GoImportPath.String()
			}

			f := g.genTestFile(gen, mData)
			outfiles = append(outfiles, f)
		}
	}
	// todo test if we can just not do this. eg return nil / empty OR return data needed for files and gen in one big batch.
	return outfiles
}

func (g *Generator) generateServiceFile(gen *protogen.Plugin, service *protogen.Service, file *protogen.File) *protogen.GeneratedFile {
	fileName := strings.ToLower(service.GoName + "/" + service.GoName + ".go") // todo format in snakecase
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{service.GoName}}.go
	f := gen.NewGeneratedFile(fileName, protogen.GoImportPath(service.GoName))

	rootGoIndent := file.GoDescriptorIdent
	pkg := getParamPKG(rootGoIndent.GoImportPath.String())
	tplate := templates.ServiceTemplate
	if g.ConnectGo {
		// in connect this could be aliased.
		pkg += "connect" // if used in this manner multiple times, tell user to make pass in correct path OR handle in templates when possible
		tplate = connectgo.ServiceTemplate
		rootGoIndent = protogen.GoIdent{GoImportPath: rootGoIndent.GoImportPath + "connect"}
	}

	if g.Firebase {
		tplate = templates.FirebaseService
	}

	s := serviceT{
		GoPkgName:   strings.ToLower(string(service.Desc.Name())),
		Imports:     rootGoIndent.GoImportPath.String(), // when called via string() it has no ""
		ServiceName: string(service.Desc.Name()),
		Pkg:         pkg,
		FullName:    string(service.Desc.FullName()),
	}

	data := templates.ExecuteTemplate(tplate, s)
	f.P(data)
	return f
}

// todo think harder about this (where should this data be kept)
type serviceT struct {
	GoPkgName   string
	ServiceName string
	Pkg         string
	FullName    string
	Imports     string
}

// FileInfo contains info from proto files needed to import generated proto to create a server
type FileInfo struct {
	Pkg           string // should be Pkg alias name.
	GoPackageName string
	GoImportPath  string
}

// collectFileData will collect file info for info required for sever generation.
func collectFileData(f *protogen.File, fileInfoMap map[string]FileInfo) map[string]FileInfo {
	goName := string(f.GoPackageName)
	if _, exists := fileInfoMap[goName]; !exists {
		fileInfoMap[goName] = FileInfo{Pkg: getParamPKG(f.GoDescriptorIdent.GoImportPath.String()), GoPackageName: string(f.GoPackageName), GoImportPath: string(f.GoImportPath)}
	}

	return fileInfoMap
}

func (g *Generator) getMethodTemplate(methodDesc protoreflect.MethodDescriptor) string {
	if !g.ConnectGo {
		return templates.MethodTemplate
	}
	switch {
	case methodDesc.IsStreamingClient() && methodDesc.IsStreamingServer():
		return connectgo.BiDirectionalStreamingTemplate
	case methodDesc.IsStreamingServer():
		return connectgo.StreamingServiceTemplate
	case methodDesc.IsStreamingClient():
		return connectgo.StreamingClientTemplate
	default:
		return connectgo.MethodTemplate
	}
}

func (g *Generator) getTestMethodTemplate(methodDesc protoreflect.MethodDescriptor) string {
	if !g.ConnectGo {
		return templates.TestFileTemplate
	}
	switch {
	case methodDesc.IsStreamingClient() && methodDesc.IsStreamingServer():
		return connectgo.TestBiDirectionalStreamFileTemplate
	case methodDesc.IsStreamingServer():
		return connectgo.TestServerStreamFileTemplate
	case methodDesc.IsStreamingClient():
		return connectgo.TestClientStreamFileTemplate
	default:
		return connectgo.TestFileTemplate
	}
}
