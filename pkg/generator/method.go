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
	ServiceName  string
	MethodName   string
	RequestType  string // input param.
	ResponseType string
	FullName     string
	Imports      string
	Pkg          string                        // proto pkg
	GoPkgName    string                        // name for pkg. Same as ServiceName but lower case.
	methodDesc   protoreflect.MethodDescriptor // for extra data from methodDescriptor.
	// for firebase trial
	ProtoPkg    string
	MessageName string
}

func (g *Generator) genRpcMethod(gen *protogen.Plugin, data methodData) *protogen.GeneratedFile {
	filename := strings.ToLower(data.ServiceName + "/" + data.MethodName + ".go")
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{method.GoName}}.go
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
			data.template = templates.FirebaseListMethod
			//data.ProtoPkg = getParamPKG(string(data.methodDesc.Output().Fields().Get(0).FullName()))
			data.ProtoPkg = data.Pkg
			//data.MessageName = getMessageNameFromPath(string(data.protoMethod.Output.Desc.Fields().Get(0).Message().FullName()))
			data.MessageName = "Example" // hard coding for now.
		}
		//data.MessageName = string(data.MethodDesc.Output().Fields().Get(0).FullName())
		//data.MethodDesc.Output().
		//data.MethodDesc.Output().Fields().Get(0).Name()
	}

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
