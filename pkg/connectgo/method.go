package connectgo

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func genRpcMethod(gen *protogen.Plugin, service *protogen.Service, method *protogen.Method) *protogen.GeneratedFile {
	filename := strings.ToLower(service.GoName + "/" + method.GoName + ".go")
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{method.GoName}}.go
	g := gen.NewGeneratedFile(filename, protogen.GoImportPath(service.GoName))

	methodCaller := genMethodCaller(service.GoName)                                // maybe methodName or methodReciever
	rootGoIndent := gen.FilesByPath[service.Location.SourceFile].GoDescriptorIdent // may run into problems depending on how files are set up.

	// todo create some func with all required pkgs imported as needed
	g.QualifiedGoIdent(rootGoIndent)
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "context", GoName: ""})
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/codes", GoName: ""})
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/status", GoName: ""})

	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "github.com/bufbuild/connect-go", GoName: ""})

	g.QualifiedGoIdent(method.Input.GoIdent)
	g.QualifiedGoIdent(method.Output.GoIdent)

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
	g := gen.NewGeneratedFile(filename, protogen.GoImportPath(service.GoName))

	g.P()
	g.P("package ", strings.ToLower(service.GoName))
	g.P()

	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "testing", GoName: ""})
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "context", GoName: ""})
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "testing", GoName: ""})
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: protogen.GoImportPath(service.GoName), GoName: ""}) // it would be nice to figure out how to have this not be aliased
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/codes", GoName: ""})        // it would be nice to figure out how to have this not be aliased
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/status", GoName: ""})       // it would be nice to figure out how to have this not be aliased
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/status", GoName: ""})       // it would be nice to figure out how to have this not be aliased
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "github.com/stretchr/testify/assert", GoName: ""})  // it would be nice to figure out how to have this not be aliased

	testFile := formatTestFile(method.GoName, service.GoName)
	g.P(testFile)
	return g
}

const methodCallerTemplate = `%s *%s`

// move to helper eventually
func genMethodCaller(in string) string {
	return fmt.Sprintf(methodCallerTemplate, strings.ToLower(in[0:1]), in)
}

// todo replace ... with gRPC method path.
const methodTemplate = `

func (%s) %s(ctx context.Context, req *connect_go.Request[%s]) (*connect_go.Response[%s], error) {
	res := connect_go.NewResponse(&%s{})
	return res, status.Error(codes.Unimplemented, "yet to be implemented")
}

`

// move to go template and use gen.Write
func formatMethod(methodCaller string, methodName string, requestType string, responseType string) string {
	return fmt.Sprintf(
		methodTemplate,
		methodCaller,
		methodName,
		requestType,
		responseType,
		responseType,
	)
}

const testFileTemplate = `
	func Test%s(t *testing.T){
		t.Parallel()
		service := &%s{}
		res, err := service.%s(context.Background(), nil)
		assert.Error(t, err)
		assert.Equal(t, codes.Unimplemented, status.Code(err))
		assert.Nil(t, res)
	}
	`

func formatTestFile(method string, service string) string {
	return fmt.Sprintf(testFileTemplate, method, service, method)
}

func getParamPKG(in string) string {
	arr := strings.Split(in, "/")
	return strings.Trim(arr[len(arr)-1], `"`)
}
