package main

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			for _, v := range f.Services {
				generateFilesForService(gen, v, f)
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

	rootGoIndent := gen.FilesByPath[service.Location.SourceFile].GoDescriptorIdent // may run into problems depending on how files are set up.
	pkg := getParamPKG(rootGoIndent.GoImportPath.String())

	structString := formatService(string(service.Desc.Name()), pkg)
	_ = g.QualifiedGoIdent(rootGoIndent) // this auto imports too.

	// todo add in template
	//t := template.Must(template.New("letter").Parse(tplate))

	g.P(structString)
	g.P()
	return g
}

func generateFilesForService(gen *protogen.Plugin, service *protogen.Service, file *protogen.File) (outfiles []*protogen.GeneratedFile) {
	serviceFile := generateServiceFile(gen, service)
	outfiles = append(outfiles, serviceFile)

	// will create a method for all services
	for _, v := range service.Methods {

		g := genRpcMethod(gen, service, v)
		outfiles = append(outfiles, g)

		// wil generate test file
		gT := genTestFile(gen, service, v)
		outfiles = append(outfiles, gT)
	}

	return outfiles
}

func genRpcMethod(gen *protogen.Plugin, service *protogen.Service, method *protogen.Method) *protogen.GeneratedFile {
	filename := strings.ToLower(service.GoName + "/" + method.GoName + ".go")
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{method.GoName}}.go
	g := gen.NewGeneratedFile(filename, protogen.GoImportPath(service.GoName))

	methodCaller := genMethodCaller(service.GoName)                                // maybe methodName or methodReciever
	rootGoIndent := gen.FilesByPath[service.Location.SourceFile].GoDescriptorIdent // may run into problems depending on how files are set up.

	// todo create some func with all required pkgs imported as needed
	g.QualifiedGoIdent(rootGoIndent)                                                                // this auto imports too.
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "context", GoName: ""})                       // it would be nice to figure out how to have this not be aliased
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/codes", GoName: ""})  // it would be nice to figure out how to have this not be aliased
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/status", GoName: ""}) // it would be nice to figure out how to have this not be aliased

	rpcfunc := formatMethod(methodCaller, method.GoName, getParamPKG(method.Input.GoIdent.GoImportPath.String())+"."+method.Input.GoIdent.GoName, getParamPKG(method.Output.GoIdent.GoImportPath.String())+"."+method.Output.GoIdent.GoName)

	g.P()
	g.P("package ", strings.ToLower(service.GoName))
	g.P()
	g.P(rpcfunc)

	return g
}

func genTestFile(gen *protogen.Plugin, service *protogen.Service, method *protogen.Method) *protogen.GeneratedFile {
	filename := strings.ToLower(service.GoName + "/" + method.GoName + "_test.go")
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{method.GoName}}.go
	gT := gen.NewGeneratedFile(filename, protogen.GoImportPath(service.GoName))

	gT.P()
	gT.P("package ", strings.ToLower(service.GoName+"_test"))
	gT.P()
	gT.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "testing", GoName: ""})
	gT.Import("github.com/stretchr/testify/assert")
	testFile := fmt.Sprintf(testFileTemplate, method.GoName)
	gT.P(testFile)
	return gT
}

func genMethodCaller(in string) string {
	return fmt.Sprintf(methodCallerTemplate, strings.ToLower(in[0:1]), in)
}

func getParamPKG(in string) string {
	arr := strings.Split(in, "/")
	return strings.Trim(arr[len(arr)-1], `"`)
}

// move to go template and use write
func formatMethod(methodCaller string, methodName string, requestType string, responseType string) string {
	return fmt.Sprintf(
		methodTemplate,
		methodName,
		methodCaller,
		methodName,
		requestType,
		responseType,
	)
}

func formatService(serviceName string, pkg string) string {
	return fmt.Sprintf(serviceTemplate, serviceName, serviceName, pkg, serviceName)
}

const serviceTemplate = `
// %s ...
type %s struct { 
	%s.Unimplemented%sServer
}
	`

const methodTemplate = `
	// %s ...
	func (%s) %s (ctx context.Context, in *%s) (out *%s, err error){
		return nil, status.Error(codes.Unimplemented, "yet to be implemented")
	}
`

const testFileTemplate = `
func Test%s(t *testing.T){
}
`

const methodCallerTemplate = `%s *%s`
