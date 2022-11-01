package generator

import (
	"fmt"
	"strings"

	"github.com/lcmaguire/protoc-gen-go-goo/pkg/connectgo"
	"github.com/lcmaguire/protoc-gen-go-goo/pkg/templates"
	"google.golang.org/protobuf/compiler/protogen"
)

func (g *Generator) generateServer(gen *protogen.Plugin, file FileInfo, services []string) {
	fileName := strings.ToLower("cmd" + "/" + file.GoPackageName + "/" + "main.go")
	f := gen.NewGeneratedFile(fileName, protogen.GoImportPath("."))

	f.P("package main ")
	imports := append(_serviceImports, protogen.GoImportPath(file.GoImportPath))

	for _, v := range imports {
		f.QualifiedGoIdent(protogen.GoIdent{GoImportPath: v})
	}

	// imports proto.
	pkg := file.Pkg
	resgisteredServices := ""
	for _, serviceName := range services {
		// dir goModPath + serviceName
		importPath := fmt.Sprintf("%s/%s", g.GoModPath, strings.ToLower(serviceName))
		f.QualifiedGoIdent(protogen.GoIdent{GoName: "", GoImportPath: protogen.GoImportPath(importPath)})

		// will probably need to be an interface or variable funcs
		if g.ConnectGo {
			resgisteredServices += templates.ExecuteTemplate(connectgo.ServiceHandleTemplate, serviceHandleData{Pkg: pkg, ServiceName: serviceName, ServiceStruct: strings.ToLower(serviceName) + "." + serviceName})
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
		f.P(templates.ExecuteTemplate(connectgo.ServerTemplate, serverData{Services: resgisteredServices}))
		return
	}
	f.P(fmt.Sprintf(templates.ServerTemplate, resgisteredServices))
}

func (g *Generator) generateConnectServer(gen *protogen.Plugin, file FileInfo, services []serviceT) {
	fileName := strings.ToLower("cmd" + "/" + file.GoPackageName + "/" + "main.go")
	f := gen.NewGeneratedFile(fileName, protogen.GoImportPath("."))

	//f.P("package main ")

	//imports := connectgo.ServiceImports
	// gets connect go gRPC.
	goPKGname := strings.ToLower(file.GoPackageName)
	connectGenImportPath := fmt.Sprintf("\"%s/%s\"", g.GoModPath, goPKGname+"connect") // template could just be {{path}}connectgo"

	pkg := file.Pkg
	resgisteredServices := ""
	servicePaths := ""
	fullNames := ""
	for _, serviceName := range services {
		// dir goModPath + serviceName
		importPath := fmt.Sprintf("%s/%s", g.GoModPath, strings.ToLower(serviceName.ServiceName))
		servicePaths += "\"" + importPath + "\"\n" // todo make func.
		fullNames += fmt.Sprintf("\"%s\",\n", serviceName.FullName)
		resgisteredServices += templates.ExecuteTemplate(connectgo.ServiceHandleTemplate, serviceHandleData{Pkg: pkg, ServiceName: serviceName.ServiceName, ServiceStruct: strings.ToLower(serviceName.ServiceName) + "." + serviceName.ServiceName})
	}

	f.P(templates.ExecuteTemplate(connectgo.ServerTemplate, serverData{Services: resgisteredServices, ConnectGenImportPath: connectGenImportPath, ServiceImports: servicePaths, FullName: fullNames}))
}

type serviceHandleData struct {
	Pkg           string
	ServiceName   string
	ServiceStruct string
}

type serverData struct {
	Services             string
	ConnectGenImportPath string
	ServiceImports       string
	FullName             string // used for reflection
}
