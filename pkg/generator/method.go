package generator

import (
	"fmt"
	"strings"

	"github.com/lcmaguire/protoc-gen-go-goo/pkg/connectgo"
	"github.com/lcmaguire/protoc-gen-go-goo/pkg/firebase"
	"github.com/lcmaguire/protoc-gen-go-goo/pkg/templates"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type methodData struct {
	MethodCaller string
	ServiceName  string
	MethodName   string
	RequestType  string
	ResponseType string
	FullName     string
	Imports      []protogen.GoIdent
	Pkg          string
	MethodDesc   protoreflect.MethodDescriptor
	MessageName  string
	protoMethod  *protogen.Method
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
		switch {
		case strings.HasPrefix(data.MethodName, "Create"):
			tplate = firebase.CreateEndpointUpdate
		case strings.HasPrefix(data.MethodName, "Update"):
			tplate = firebase.UpdateEndpointTemplate
		case strings.HasPrefix(data.MethodName, "Delete"):
			tplate = firebase.DeleteEndpointTemplate
		case strings.HasPrefix(data.MethodName, "Get"):
			tplate = firebase.GetEndpointTemplate
		case strings.HasPrefix(data.MethodName, "List"):
			tplate = firebase.ListEndpointTemplate
			data.MessageName = string(data.protoMethod.Output.Desc.Fields().Get(0).Name())
			//data.MessageName = string(data.MethodDesc.Output().Fields().Get(0).FullName())
			//data.MethodDesc.Output().
			//data.MethodDesc.Output().Fields().Get(0).Name()
		}

		switch {
		case data.MethodDesc.IsStreamingClient() && data.MethodDesc.IsStreamingServer():
			imports = append(imports, "errors", "io", "fmt")
			tplate = connectgo.BiDirectionalStreamingTemplate
		case data.MethodDesc.IsStreamingServer():
			imports = append(imports, "time")
			tplate = connectgo.StreamingServiceTemplate
		case data.MethodDesc.IsStreamingClient():
			imports = append(imports, "errors")
			tplate = connectgo.StreamingClientTemplate
		}
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
		switch {
		case data.MethodDesc.IsStreamingClient() && data.MethodDesc.IsStreamingServer():
			tplate = connectgo.TestBiDirectionalStreamFileTemplate
			imports = connectgo.TestBiDirectionalMethod
			// imports connect go gRPC.
			connectGenImportPath := fmt.Sprintf("%s/%s", g.GoModPath, data.Pkg+"connect")
			f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: protogen.GoImportPath(connectGenImportPath)})
		case data.MethodDesc.IsStreamingClient():
			tplate = connectgo.TestClientStreamFileTemplate
			imports = connectgo.TestClientStreamMethod
			connectGenImportPath := fmt.Sprintf("%s/%s", g.GoModPath, data.Pkg+"connect")
			f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: protogen.GoImportPath(connectGenImportPath)})
		case data.MethodDesc.IsStreamingServer():
			tplate = connectgo.TestServerStreamFileTemplate
			imports = connectgo.TestServerStreamMethod
			connectGenImportPath := fmt.Sprintf("%s/%s", g.GoModPath, data.Pkg+"connect")
			f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: protogen.GoImportPath(connectGenImportPath)})
		}
	}

	// move below to funcs.
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
