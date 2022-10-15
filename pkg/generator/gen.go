package generator

import (
	"fmt"
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

// todo gen constructor.
func (g *Generator) generateServiceFile(gen *protogen.Plugin, service *protogen.Service) *protogen.GeneratedFile { // consider returning []
	fileName := strings.ToLower(service.GoName + "/" + service.GoName + ".go") // todo format in snakecase
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{service.GoName}}.go
	f := gen.NewGeneratedFile(fileName, protogen.GoImportPath(service.GoName))
	f.P()
	f.P("package ", strings.ToLower(service.GoName))
	f.P()

	rootGoIndent := gen.FilesByPath[service.Location.SourceFile].GoDescriptorIdent // may run into problems depending on how files are set up.
	pkg := getParamPKG(rootGoIndent.GoImportPath.String())

	structString := formatService(string(service.Desc.Name()), pkg)
	_ = f.QualifiedGoIdent(rootGoIndent) // this auto imports too.

	// todo add in template
	//t := template.Must(template.New("letter").Parse(tplate))

	f.P(structString)
	f.P()
	return f
}

func (g *Generator) generateFilesForService(gen *protogen.Plugin, service *protogen.Service, file *protogen.File) (outfiles []*protogen.GeneratedFile) {
	serviceFile := g.generateServiceFile(gen, service)
	outfiles = append(outfiles, serviceFile)

	// will create a method for all services
	for _, v := range service.Methods {

		f := g.genRpcMethod(gen, service, v)
		outfiles = append(outfiles, f)

	}

	if g.Tests {
		// i wonder if time complexity is important for proto plugins.
		// and diff between looping twice vs looping once with an if statement is, ill use my brain later and figure it out
		for _, v := range service.Methods {
			f := g.genTestFile(gen, service, v)
			outfiles = append(outfiles, f)
		}
	}

	// todo test if we can just not do this. eg return nil / empty
	return outfiles
}

func (g *Generator) genRpcMethod(gen *protogen.Plugin, service *protogen.Service, method *protogen.Method) *protogen.GeneratedFile {
	filename := strings.ToLower(service.GoName + "/" + method.GoName + ".go")
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{method.GoName}}.go
	f := gen.NewGeneratedFile(filename, protogen.GoImportPath(service.GoName))

	methodCaller := genMethodCaller(service.GoName)                                // maybe methodName or methodReciever
	rootGoIndent := gen.FilesByPath[service.Location.SourceFile].GoDescriptorIdent // may run into problems depending on how files are set up.

	// todo create some func with all required pkgs imported as needed
	f.QualifiedGoIdent(rootGoIndent)                                                                // this auto imports too.
	f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "context", GoName: ""})                       // it would be nice to figure out how to have this not be aliased
	f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/codes", GoName: ""})  // it would be nice to figure out how to have this not be aliased
	f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/status", GoName: ""}) // it would be nice to figure out how to have this not be aliased

	f.QualifiedGoIdent(method.Input.GoIdent)
	f.QualifiedGoIdent(method.Output.GoIdent)

	rpcfunc := formatMethod(methodCaller, method.GoName, getParamPKG(method.Input.GoIdent.GoImportPath.String())+"."+method.Input.GoIdent.GoName, getParamPKG(method.Output.GoIdent.GoImportPath.String())+"."+method.Output.GoIdent.GoName)

	f.P()
	f.P("package ", strings.ToLower(service.GoName))
	f.P()
	f.P(rpcfunc)

	return f
}

func (g *Generator) genTestFile(gen *protogen.Plugin, service *protogen.Service, method *protogen.Method) *protogen.GeneratedFile {
	filename := strings.ToLower(service.GoName + "/" + method.GoName + "_test.go")
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{method.GoName}}.go
	f := gen.NewGeneratedFile(filename, protogen.GoImportPath(service.GoName))

	f.P()
	f.P("package ", strings.ToLower(service.GoName))
	f.P()

	f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "testing", GoName: ""})
	f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "context", GoName: ""})
	f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "testing", GoName: ""})
	f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: protogen.GoImportPath(service.GoName), GoName: ""}) // it would be nice to figure out how to have this not be aliased
	f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/codes", GoName: ""})        // it would be nice to figure out how to have this not be aliased
	f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/status", GoName: ""})       // it would be nice to figure out how to have this not be aliased
	f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/status", GoName: ""})       // it would be nice to figure out how to have this not be aliased
	f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "github.com/stretchr/testify/assert", GoName: ""})  // it would be nice to figure out how to have this not be aliased

	testFile := formatTestFile(method.GoName, service.GoName)
	f.P(testFile)
	return f
}

// need pkg, services,
func (g *Generator) generateServer(gen *protogen.Plugin, file *protogen.File) {
	services := []string{}

	for _, v := range file.Services {
		services = append(services, v.GoName) // service.Desc.Name()
	}

	fileName := strings.ToLower("cmd" + "/" + string(file.GoPackageName) + "/" + "main.go")
	f := gen.NewGeneratedFile(fileName, protogen.GoImportPath("."))

	f.P("package main ")

	// required imports
	f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: file.GoImportPath})
	f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "log"})
	f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "net"})
	f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc"})
	f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/reflection"})

	// hard coding these vals for now, will have to think of a cleaner way of figuring out go mod path + out path for generated services.
	const _hardCodedPath = "github.com/lcmaguire/protoc-gen-go-goo"
	const _hardCodedGoGooOutPath = "example"

	pkg := getParamPKG(file.GoDescriptorIdent.GoImportPath.String())

	resgisteredServices := ""
	for _, serviceName := range services {
		// dir goModPath + serviceName
		// will also need to get go-goo_out path to put inbetween
		importPath := fmt.Sprintf("%s/%s/%s", _hardCodedPath, _hardCodedGoGooOutPath, strings.ToLower(serviceName))
		f.QualifiedGoIdent(protogen.GoIdent{GoName: "", GoImportPath: protogen.GoImportPath(importPath)})

		resgisteredServices += fmt.Sprintf(
			templates.RegisterServiceTemplate,
			pkg,
			serviceName,
			strings.ToLower(serviceName)+"."+serviceName,
		)
	}

	f.P(fmt.Sprintf(templates.ServerTemplate, resgisteredServices))
}
