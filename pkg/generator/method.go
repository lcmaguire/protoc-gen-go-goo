package generator

import (
	"fmt"
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
	Pkg          string
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
		switch {
		case data.methodDesc.IsStreamingClient() && data.methodDesc.IsStreamingServer():
			imports = append(imports, "errors", "io", "fmt")
			tplate = connectgo.BiDirectionalStreamingTemplate
		case data.methodDesc.IsStreamingServer():
			imports = append(imports, "time")
			tplate = connectgo.ServerStreamingTemplate
		case data.methodDesc.IsStreamingClient():
			imports = append(imports, "errors")
			tplate = connectgo.ClientStreamingTemplate
		}
	} else {
		switch {
		case data.methodDesc.IsStreamingClient() && data.methodDesc.IsStreamingServer():
			tplate = templates.BidirectionStreamingMethod
		case data.methodDesc.IsStreamingServer():
			tplate = templates.ServerStreamingMethod
		case data.methodDesc.IsStreamingClient():
			tplate = templates.ClientStreamingMethod
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
	}

	switch {
	case data.methodDesc.IsStreamingClient() && data.methodDesc.IsStreamingServer():
		tplate = connectgo.TestBidirectionalStreamTemplate // streaming types are hard to make tests for.
		imports = connectgo.TestBiDirectionalMethod
		// imports connect go gRPC.
		connectGenImportPath := fmt.Sprintf("%s/%s", g.GoModPath, data.Pkg+"connect")
		f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: protogen.GoImportPath(connectGenImportPath)})
	case data.methodDesc.IsStreamingClient():
		tplate = connectgo.TestClientStreamTemplate // streaming types are hard to make tests for.
		imports = connectgo.TestClientStreamMethod  // make seperate for clientStreaming.
		connectGenImportPath := fmt.Sprintf("%s/%s", g.GoModPath, data.Pkg+"connect")
		f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: protogen.GoImportPath(connectGenImportPath)})
	case data.methodDesc.IsStreamingServer():
		tplate = connectgo.UnsportedTestFile // streaming types are hard to make tests for.
		imports = []protogen.GoImportPath{"testing", "github.com/stretchr/testify/assert"}
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
