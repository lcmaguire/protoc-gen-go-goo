package generator

import (
	"html/template"
	"os"
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

	type serviceT struct {
		ServiceName string
		Pkg         string
	}

	s := serviceT{
		ServiceName: string(service.Desc.Name()),
		Pkg:         pkg,
	}
	// todo see what this means.

	// could try f.Bytes + x.
	templ, err := template.New("test").Parse(templates.ActualServiceTemplate)
	if err != nil {
		return nil
	}

	if err := templ.Execute(os.Stdout, s); err != nil {
		panic(err)
	}

	// todo add in template
	//t := template.Must(template.New("letter").Parse(tplate))

	f.P(structString)
	f.P()
	return f
}
