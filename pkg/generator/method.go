package generator

import (
	"fmt"
	"strings"

	"github.com/lcmaguire/protoc-gen-go-goo/pkg/connectgo"
	connectgotemplates "github.com/lcmaguire/protoc-gen-go-goo/pkg/templates/connecttemplates"
	"google.golang.org/protobuf/compiler/protogen"
)

func (g *Generator) genRpcMethod(gen *protogen.Plugin, service *protogen.Service, method *protogen.Method) *protogen.GeneratedFile {
	filename := strings.ToLower(service.GoName + "/" + method.GoName + ".go")
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{method.GoName}}.go
	f := gen.NewGeneratedFile(filename, protogen.GoImportPath(service.GoName))

	methodCaller := genMethodCaller(service.GoName)                                // maybe methodName or methodReciever
	rootGoIndent := gen.FilesByPath[service.Location.SourceFile].GoDescriptorIdent // may run into problems depending on how files are set up.

	// todo create some func with all required pkgs imported as needed
	imports := _methodImports
	if g.ConnectGo {
		imports = connectgo.MethodImports
	}

	f.QualifiedGoIdent(rootGoIndent) // this auto imports too.
	for _, v := range imports {      // should be func
		f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: v})
	}
	f.QualifiedGoIdent(method.Input.GoIdent)
	f.QualifiedGoIdent(method.Output.GoIdent)

	request := getParamPKG(method.Input.GoIdent.GoImportPath.String()) + "." + method.Input.GoIdent.GoName
	response := getParamPKG(method.Output.GoIdent.GoImportPath.String()) + "." + method.Output.GoIdent.GoName
	rpcfunc := formatMethod(methodCaller, method.GoName, request, response)
	if g.ConnectGo {
		rpcfunc = fmt.Sprintf(connectgotemplates.MethodTemplate,
			methodCaller,
			method.GoName,
			request,
			response,
			response)
	}

	f.P()
	f.P("package ", strings.ToLower(service.GoName))
	f.P()
	f.P(rpcfunc)

	return f
}
