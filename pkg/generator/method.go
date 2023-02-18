package generator

import (
	"strings"

	"github.com/lcmaguire/protoc-gen-go-goo/pkg/templates"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// todo figure our what can be removed from this struct
// methodData contains data for generating Methods via a template.
type methodData struct {
	template     string
	testTemplate string
	MethodCaller string
	ServiceChar  string
	ServiceName  string
	MethodName   string
	RequestType  string // input param.
	ResponseType string
	FullName     string
	Imports      string
	Pkg          string                        // proto pkg
	GoPkgName    string                        // name for pkg. Same as ServiceName but lower case.
	methodDesc   protoreflect.MethodDescriptor // for extra data from methodDescriptor.
}

func (g *Generator) genRpcMethod(gen *protogen.Plugin, data methodData) *protogen.GeneratedFile {
	filename := strings.ToLower(data.ServiceName + "/" + data.MethodName + ".go")
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{method.GoName}}.go
	f := gen.NewGeneratedFile(filename, protogen.GoImportPath(data.ServiceName))
	rpcfunc := templates.ExecuteTemplate(data.template, data)
	f.P(rpcfunc)
	return f
}

func (g *Generator) genTestFile(gen *protogen.Plugin, data methodData) *protogen.GeneratedFile {
	filename := strings.ToLower(data.ServiceName + "/" + data.MethodName + "_test.go")
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{method.GoName}}.go
	f := gen.NewGeneratedFile(filename, protogen.GoImportPath(data.ServiceName))
	testFile := templates.ExecuteTemplate(data.testTemplate, data)
	f.P(testFile)
	return f
}
