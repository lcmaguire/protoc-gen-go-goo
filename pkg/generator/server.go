package generator

import (
	"fmt"
	"strings"

	"github.com/lcmaguire/protoc-gen-go-goo/pkg/connectgo"
	"github.com/lcmaguire/protoc-gen-go-goo/pkg/templates"
	"google.golang.org/protobuf/compiler/protogen"
)

// need pkg, services,
func (g *Generator) generateServer(gen *protogen.Plugin, file *protogen.File, services []string) {
	fileName := strings.ToLower("cmd" + "/" + string(file.GoPackageName) + "/" + "main.go")
	f := gen.NewGeneratedFile(fileName, protogen.GoImportPath("."))

	f.P("package main ")
	imports := append(_serviceImports, protogen.GoImportPath(file.GoImportPath))
	if g.ConnectGo {
		imports = connectgo.ServiceImports
		goPKGname := strings.ToLower(string(file.GoPackageName))
		connectGenImportPath := fmt.Sprintf("%s/%s", g.GoModPath, goPKGname+"connect")
		f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: protogen.GoImportPath(connectGenImportPath)})
	}

	for _, v := range imports {
		f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: v})
	}

	pkg := getParamPKG(file.GoDescriptorIdent.GoImportPath.String())

	resgisteredServices := ""
	for _, serviceName := range services {
		// dir goModPath + serviceName
		importPath := fmt.Sprintf("%s/%s", g.GoModPath, strings.ToLower(serviceName))
		if g.ConnectGo {
			importPath = fmt.Sprintf("%s/%s", g.GoModPath, strings.ToLower(serviceName))
		}
		f.QualifiedGoIdent(protogen.GoIdent{GoName: "", GoImportPath: protogen.GoImportPath(importPath)})

		// will probably need to be an interface or variable funcs
		if g.ConnectGo {
			resgisteredServices += fmt.Sprintf(
				// todo move to template
				connectgo.ServiceHandleTemplate,
				pkg,
				serviceName,
				strings.ToLower(serviceName)+"."+serviceName,
			)
		} else {
			resgisteredServices += fmt.Sprintf(
				templates.RegisterServiceTemplate,
				pkg,
				serviceName,
				strings.ToLower(serviceName)+"."+serviceName,
			)
		}
	}

	if g.ConnectGo {
		f.P(fmt.Sprintf(connectgo.ServerTemplate, resgisteredServices))
		return
	}
	f.P(fmt.Sprintf(templates.ServerTemplate, resgisteredServices))
}
