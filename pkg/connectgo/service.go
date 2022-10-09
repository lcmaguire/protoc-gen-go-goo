package connectgo

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func generateServiceFile(gen *protogen.Plugin, service *protogen.Service) *protogen.GeneratedFile { // consider returning []
	fileName := strings.ToLower(service.GoName + "/" + service.GoName + ".go") // todo format in snakecase
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{service.GoName}}.go
	g := gen.NewGeneratedFile(fileName, protogen.GoImportPath(service.GoName))
	g.P()
	g.P("package ", strings.ToLower(service.GoName))
	g.P()
	g.P("// TSUUUU")

	rootGoIndent := gen.FilesByPath[service.Location.SourceFile].GoDescriptorIdent // may run into problems depending on how files are set up.
	arr := strings.Split(rootGoIndent.GoImportPath.String(), "/")
	pkg := strings.Trim(arr[len(arr)-1], `"`)

	g.QualifiedGoIdent(rootGoIndent)
	structString := formatService(string(service.Desc.Name()), pkg)

	g.P(structString)
	g.P()
	return g
}

const serviceTemplate = `
// %s ...
type %s struct { 
	%s.Unimplemented%sServer
}
	`

func formatService(serviceName string, pkg string) string {
	return fmt.Sprintf(serviceTemplate, serviceName, serviceName, pkg, serviceName)
}
