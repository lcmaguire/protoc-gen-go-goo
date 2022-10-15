package generator

import (
	"fmt"
	"strings"

	"github.com/lcmaguire/protoc-gen-go-goo/pkg/connectgo"
	"github.com/lcmaguire/protoc-gen-go-goo/pkg/templates"
	connectgotemplates "github.com/lcmaguire/protoc-gen-go-goo/pkg/templates/connecttemplates"
	"google.golang.org/protobuf/compiler/protogen"
)

// need pkg, services,
func (g *Generator) generateServer(gen *protogen.Plugin, file *protogen.File) {
	services := []string{}

	for _, v := range file.Services {
		services = append(services, v.GoName) // service.Desc.Name()
	}

	fileName := strings.ToLower("cmd" + "/" + string(file.GoPackageName) + "/" + "main.go")
	f := gen.NewGeneratedFile(fileName, protogen.GoImportPath("."))

	f.P("package main ")
	// hardcoding for now
	const _hardCodedPath = "github.com/lcmaguire/protoc-gen-go-goo/example" // if connect + connect

	imports := _serviceImports
	if g.ConnectGo {
		// get different imports
		imports = connectgo.ServiceImports
		goPKGname := strings.ToLower(string(file.GoPackageName))
		connectGenImportPath := fmt.Sprintf("%s/%s", _hardCodedPath+"connect", goPKGname+"connect")
		f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: protogen.GoImportPath(connectGenImportPath)})

	}

	imports = append(imports, protogen.GoImportPath(file.GoImportPath))
	for _, v := range imports {
		f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: v}) // service imports should probably be GoIdent
	}

	pkg := getParamPKG(file.GoDescriptorIdent.GoImportPath.String())
	// try strings.ToLower(string(file.GoPackageName))

	resgisteredServices := ""
	for _, serviceName := range services {
		// dir goModPath + serviceName
		importPath := fmt.Sprintf("%s/%s", _hardCodedPath, strings.ToLower(serviceName))
		f.QualifiedGoIdent(protogen.GoIdent{GoName: "", GoImportPath: protogen.GoImportPath(importPath)})

		// will probably need to be an interface or variable funcs
		if g.ConnectGo {
			resgisteredServices += fmt.Sprintf(
				connectgotemplates.ServiceHandleTemplate,
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
		f.P(fmt.Sprintf(connectgotemplates.ConnectGoServerTemplate, resgisteredServices))
		return
	}
	f.P(fmt.Sprintf(templates.ServerTemplate, resgisteredServices))
}
