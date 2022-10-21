package generator

import (
	"strings"

	"github.com/lcmaguire/protoc-gen-go-goo/pkg/connectgo"
	"github.com/lcmaguire/protoc-gen-go-goo/pkg/templates"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type methodData struct {
	S1           string // its probably better to have method caller
	ServiceName  string
	MethodName   string
	RequestType  string
	ResponseType string
	FullName     string
	Imports      []protogen.GoIdent
	methodDesc   protoreflect.MethodDescriptor
}

func (g *Generator) genRpcMethod(gen *protogen.Plugin, data methodData) *protogen.GeneratedFile {
	filename := strings.ToLower(data.ServiceName + "/" + data.MethodName + ".go")
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{method.GoName}}.go
	f := gen.NewGeneratedFile(filename, protogen.GoImportPath(data.ServiceName))

	imports := _methodImports
	tplate := templates.MethodTemplate
	if g.ConnectGo {
		imports = connectgo.MethodImports
		tplate = connectgo.MethodTemplate
	}

	if data.methodDesc.IsStreamingClient() {
		imports = append(imports, "errors")
		tplate = connectgo.ClientStreamingTemplate
	}

	if data.methodDesc.IsStreamingServer() {
		//imports = append(imports, "errors", "fmt")
		tplate = connectgo.ServerStreamingTemplate
	}

	if data.methodDesc.IsStreamingClient() && data.methodDesc.IsStreamingServer() {
		imports = append(imports, "errors", "io", "fmt")
		tplate = connectgo.BiDirectionalStreamingTemplate
	}

	// these are always imported.
	for _, v := range data.Imports {
		f.QualifiedGoIdent(v)
	}
	for _, v := range imports {
		f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: v})
	}

	rpcfunc := templates.ExecuteTemplate(tplate, data)

	f.P()
	f.P("package ", strings.ToLower(data.ServiceName))
	f.P()
	f.P(rpcfunc)

	return f
}

func (g *Generator) genTestFile(gen *protogen.Plugin, data methodData) *protogen.GeneratedFile {
	filename := strings.ToLower(data.ServiceName + "/" + data.MethodName + "_test.go")
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{method.GoName}}.go
	f := gen.NewGeneratedFile(filename, protogen.GoImportPath(data.ServiceName))

	f.P()
	f.P("package ", strings.ToLower(data.ServiceName))
	f.P()

	imports := _testImports
	tplate := templates.TestFileTemplate
	if g.ConnectGo {
		imports = connectgo.TestImports
		tplate = connectgo.TestFileTemplate
	}
	for _, v := range data.Imports {
		f.QualifiedGoIdent(v)
	}
	for _, v := range imports {
		f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: v})
	}
	testFile := templates.ExecuteTemplate(tplate, data)
	f.P(testFile)
	return f
}
