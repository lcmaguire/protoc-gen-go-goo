package generator

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func (g *Generator) genTestFile(gen *protogen.Plugin, service *protogen.Service, method *protogen.Method) *protogen.GeneratedFile {
	filename := strings.ToLower(service.GoName + "/" + method.GoName + "_test.go")
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{method.GoName}}.go
	f := gen.NewGeneratedFile(filename, protogen.GoImportPath(service.GoName))

	f.P()
	f.P("package ", strings.ToLower(service.GoName))
	f.P()

	f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: protogen.GoImportPath(service.GoName), GoName: ""})
	for _, v := range _testImports {
		f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: v})
	}

	if g.ConnectGo {
		// import connect go code + format in + out to make sense

	}

	testFile := formatTestFile(method.GoName, service.GoName)
	f.P(testFile)
	return f
}
