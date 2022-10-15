package generator

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func (g *Generator) genRpcMethod(gen *protogen.Plugin, service *protogen.Service, method *protogen.Method) *protogen.GeneratedFile {
	filename := strings.ToLower(service.GoName + "/" + method.GoName + ".go")
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{method.GoName}}.go
	f := gen.NewGeneratedFile(filename, protogen.GoImportPath(service.GoName))

	methodCaller := genMethodCaller(service.GoName)                                // maybe methodName or methodReciever
	rootGoIndent := gen.FilesByPath[service.Location.SourceFile].GoDescriptorIdent // may run into problems depending on how files are set up.

	// todo create some func with all required pkgs imported as needed
	f.QualifiedGoIdent(rootGoIndent)   // this auto imports too.
	for _, v := range _methodImports { // should be func
		f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: v})
	}
	f.QualifiedGoIdent(method.Input.GoIdent)
	f.QualifiedGoIdent(method.Output.GoIdent)

	rpcfunc := formatMethod(methodCaller, method.GoName, getParamPKG(method.Input.GoIdent.GoImportPath.String())+"."+method.Input.GoIdent.GoName, getParamPKG(method.Output.GoIdent.GoImportPath.String())+"."+method.Output.GoIdent.GoName)

	f.P()
	f.P("package ", strings.ToLower(service.GoName))
	f.P()
	f.P(rpcfunc)

	return f
}
