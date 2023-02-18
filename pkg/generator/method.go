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
	GoPkgName    string                        // PKG for where service will go. Same as ServiceName but lower case.
	methodDesc   protoreflect.MethodDescriptor // for extra data from methodDescriptor.

	// Workaround for List for firebase option
	MessageName string // the name of the Message that will be used by the service. e.g. MyMessage
}

func (g *Generator) genRpcMethod(gen *protogen.Plugin, data methodData) *protogen.GeneratedFile {
	filename := strings.ToLower(data.ServiceName + "/" + data.MethodName + ".go")

	// TOOD have this be some kind of func in the GeneratorStruct
	if g.Firebase {
		switch {
		case strings.HasPrefix(data.MethodName, "Create"):
			data.template = templates.FirebaseCreateMethod
		case strings.HasPrefix(data.MethodName, "Update"):
			data.template = templates.FirebaseUpdateMethod
		case strings.HasPrefix(data.MethodName, "Delete"):
			data.template = templates.FirebaseDeleteMethod
		case strings.HasPrefix(data.MethodName, "Get"):
			data.template = templates.FirebaseGetMethod
		case strings.HasPrefix(data.MethodName, "List"):
			// TODO look at having below be done cleaner ( perhaps via annotations).
			a := data.methodDesc.Output().Name()
			b := strings.TrimPrefix(string(a), "List")
			b = strings.TrimSuffix(b, "Response")
			data.MessageName = b

			data.template = templates.FirebaseListMethod
		}
	}

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
