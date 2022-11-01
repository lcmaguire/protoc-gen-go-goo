package generator

import (
	"fmt"
	"strings"

	"github.com/lcmaguire/protoc-gen-go-goo/pkg/templates"
	"google.golang.org/protobuf/compiler/protogen"
)

/*
func (g *Generator) generateServer(gen *protogen.Plugin, file FileInfo, services []string) {

	fileName := strings.ToLower("cmd" + "/" + file.GoPackageName + "/" + "main.go")
	f := gen.NewGeneratedFile(fileName, protogen.GoImportPath("."))

	// f.P("package main ")
	serverData := serverData{
		GenImportPath: g.GoModPath,
	}
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
		resgisteredServices += fmt.Sprintf(
			templates.RegisterServiceTemplate,
			pkg,
			serviceName,
			strings.ToLower(serviceName)+"."+serviceName,
		)

	}

	data := templates.ExecuteTemplate(connectgo.ServerTemplate, serverData)
	f.P(data)
}
*/

func (g *Generator) generateConnectServer(gen *protogen.Plugin, file FileInfo, services []serviceT) {
	fileName := strings.ToLower("cmd" + "/" + file.GoPackageName + "/" + "main.go")
	f := gen.NewGeneratedFile(fileName, protogen.GoImportPath("."))

	//f.P("package main ")

	//imports := connectgo.ServiceImports
	// gets connect go gRPC.
	genCodeImportPath := g.GoModPath
	if g.ConnectGo { // maybe this should be handled prior or by the incoming g.GoModPath or template.
		goPKGname := strings.ToLower(file.GoPackageName)
		genCodeImportPath = fmt.Sprintf("%s/%s", g.GoModPath, goPKGname)
	}

	pkg := file.Pkg
	resgisteredServices := ""
	servicePaths := ""
	fullNames := ""
	for _, serviceName := range services {
		// dir goModPath + serviceName
		importPath := fmt.Sprintf("%s/%s", g.GoModPath, strings.ToLower(serviceName.ServiceName))
		servicePaths += "\"" + importPath + "\"\n" // todo make func.
		fullNames += fmt.Sprintf("\"%s\",\n", serviceName.FullName)
		resgisteredServices += templates.ExecuteTemplate(g.RegisterServerTemplate, serviceHandleData{Pkg: pkg, ServiceName: serviceName.ServiceName, ServiceStruct: strings.ToLower(serviceName.ServiceName) + "." + serviceName.ServiceName})
	}

	f.P(templates.ExecuteTemplate(g.ServerTemplate, serverData{Services: resgisteredServices, GenImportPath: genCodeImportPath, ServiceImports: servicePaths, FullName: fullNames}))
}

// serverData for registering specific services
type serviceHandleData struct {
	Pkg           string
	ServiceName   string
	ServiceStruct string
}

// serverData for the server you will be generating.
type serverData struct {
	Services       string // the services being registered.
	GenImportPath  string // import path for the service.
	ServiceImports string // what is imported by the func
	FullName       string // used for reflection
}
