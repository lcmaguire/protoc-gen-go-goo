package main

import (
	"flag"
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

var root *string //= flag.String("root", "", "")

func main() {
	var flags flag.FlagSet
	root = flags.String("root", "", "")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			for _, v := range f.Services {
				generateFilesForService(gen, v)
			}
			// try using protoreflect.
		}
		return nil
	})
}

// generateFile generates a _ascii.pb.go file containing gRPC service definitions.
func generateServiceFile(gen *protogen.Plugin, service *protogen.Service) *protogen.GeneratedFile { // consider returning []
	fileName := strings.ToLower(service.GoName + "/" + service.GoName + ".go") // todo format in snakecase
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{service.GoName}}.go
	g := gen.NewGeneratedFile(fileName, protogen.GoImportPath(service.GoName))
	g.P()
	g.P("package ", strings.ToLower(service.GoName))
	g.P()

	rootGoIndent := rootPathGoIndent(gen.FilesByPath[service.Location.SourceFile]) // may run into problems depending on how files are set up.

	structString := fmt.Sprintf(tplate, service.Desc.Name(), rootGoIndent.GoName, service.Desc.Name())
	_ = g.QualifiedGoIdent(rootGoIndent) // this auto imports too.

	// todo add in template
	//t := template.Must(template.New("letter").Parse(tplate))

	g.P(structString)
	g.P()
	return g
}

func generateFilesForService(gen *protogen.Plugin, service *protogen.Service) (outfiles []*protogen.GeneratedFile) {
	serviceFile := generateServiceFile(gen, service)
	outfiles = append(outfiles, serviceFile)

	// will create a method for all services
	for _, v := range service.Methods {
		filename := strings.ToLower(service.GoName + "/" + v.GoName + ".go")
		// will be in format /{{goo_out_path}}/{{service.GoName}}/{{method.GoName}}.go
		g := gen.NewGeneratedFile(filename, protogen.GoImportPath(service.GoName))

		methodCaller := genMethodCaller(service.GoName)                                // maybe methodName or methodReciever
		rootGoIndent := rootPathGoIndent(gen.FilesByPath[service.Location.SourceFile]) // may run into problems depending on how files are set up.

		g.QualifiedGoIdent(rootGoIndent) // this auto imports too.
		// todo create some func with all required pkgs imported as needed
		g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "context", GoName: ""}) // it would be nice to figure out how to have this not be aliased

		rpcfunc := fmt.Sprintf(
			methodTemplate,
			methodCaller,
			v.GoName,
			rootGoIndent.GoName+"."+v.Input.GoIdent.GoName,
			rootGoIndent.GoName+"."+v.Output.GoIdent.GoName,
		)

		g.P()
		g.P("package ", strings.ToLower(service.GoName))
		g.P()
		g.P(rpcfunc)

		outfiles = append(outfiles, g)

		filename = strings.ToLower(service.GoName + "/" + v.GoName + "_test.go")
		// will be in format /{{goo_out_path}}/{{service.GoName}}/{{method.GoName}}.go
		gT := gen.NewGeneratedFile(filename, protogen.GoImportPath(service.GoName))

		gT.P()
		gT.P("package ", strings.ToLower(service.GoName+"_test"))
		gT.P()
		gT.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "testing", GoName: ""})
		gT.Import("github.com/stretchr/testify/assert")
		testFile := fmt.Sprintf(testFileTemplate, service.GoName)
		gT.P(testFile)

		outfiles = append(outfiles, gT)
	}

	return outfiles
}

// todo move to using templates or something nicer.
var tplate = ` type %s struct 
{ 
	%s.Unimplemented%sServer
}
	`

var methodTemplate = `
	func (%s) %s (ctx context.Context, in *%s) (out *%s, err error){
		return
	}
`

var testFileTemplate = `
func Test%s(t *testing.T){
}
`

var methodCallerTemplate = `%s *%s`

func genMethodCaller(in string) string {
	return fmt.Sprintf(methodCallerTemplate, strings.ToLower(in[0:1]), in)
}

func rootPathGoIndent(f *protogen.File) protogen.GoIdent {
	importPath := f.GoImportPath
	pkgName := f.GoPackageName
	path := fmt.Sprintf("%s/%s", *root, string(importPath))
	return protogen.GoImportPath(path).Ident(string(pkgName))
}
