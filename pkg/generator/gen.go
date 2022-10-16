package generator

import (
	"strings"

	"github.com/lcmaguire/protoc-gen-go-goo/pkg/templates"
	"google.golang.org/protobuf/compiler/protogen"
)

/*
	What I want
	Struct that has all info from protogen.
	funcs that handle all string manipulation / contatonation
	options that allow for certain things to happen in certain orders.

	if connect == true.

	then when GenerateMethod It SHOULD generate a connect compatable thing.

*/

// Generator holds all info for generating boiler plate code.
// consider this being purely cfg and creating another more useful struct
type Generator struct {
	ConnectGo bool // either used as bool to decide template, or an interface for different generation. For now bool (maybe load templates based upon this)
	Server    bool
	GoModPath string
	Tests     bool
}

func (g *Generator) Run(gen *protogen.Plugin) error {
	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}

		for _, v := range f.Services {
			g.generateFilesForService(gen, v, f)
		}

		if g.Server {
			g.generateServer(gen, f)
		}
	}
	return nil
}

func (g *Generator) generateFilesForService(gen *protogen.Plugin, service *protogen.Service, file *protogen.File) (outfiles []*protogen.GeneratedFile) {
	serviceFile := g.generateServiceFile(gen, service)
	outfiles = append(outfiles, serviceFile)

	// will create a method for all services
	for _, v := range service.Methods {
		mData := methodData{
			S1:           strings.ToLower(service.GoName[0:1]),
			ServiceName:  service.GoName,
			MethodName:   v.GoName,
			FullName:     string(v.Desc.FullName()),
			RequestType:  getParamPKG(v.Input.GoIdent.GoImportPath.String()) + "." + v.Input.GoIdent.GoName,
			ResponseType: getParamPKG(v.Output.GoIdent.GoImportPath.String()) + "." + v.Output.GoIdent.GoName,
			Imports:      []protogen.GoIdent{v.Input.GoIdent, v.Output.GoIdent, {GoImportPath: protogen.GoImportPath(service.GoName)}},
		}
		f := g.genRpcMethod(gen, mData)
		outfiles = append(outfiles, f)
		if g.Tests {
			f := g.genTestFile(gen, mData)
			outfiles = append(outfiles, f)
		}
	}
	// todo test if we can just not do this. eg return nil / empty
	return outfiles
}

func (g *Generator) generateServiceFile(gen *protogen.Plugin, service *protogen.Service) *protogen.GeneratedFile { // consider returning []
	fileName := strings.ToLower(service.GoName + "/" + service.GoName + ".go") // todo format in snakecase
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{service.GoName}}.go
	f := gen.NewGeneratedFile(fileName, protogen.GoImportPath(service.GoName))
	f.P()
	f.P("package ", strings.ToLower(service.GoName))
	f.P()

	rootGoIndent := gen.FilesByPath[service.Location.SourceFile].GoDescriptorIdent // may run into problems depending on how files are set up.
	pkg := getParamPKG(rootGoIndent.GoImportPath.String())
	_ = f.QualifiedGoIdent(rootGoIndent)

	// todo think harder about this (where should this data be kept)
	type serviceT struct {
		ServiceName string
		Pkg         string
		FullName    string
	}
	s := serviceT{
		ServiceName: string(service.Desc.Name()),
		Pkg:         pkg,
		FullName:    string(service.Desc.FullName()),
	}

	data := templates.ExecuteTemplate(templates.ServiceTemplate, s)
	f.P(data)
	f.P()
	return f
}
