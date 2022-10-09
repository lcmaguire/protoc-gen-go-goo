package connectgo

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

// sampled from https://connect.build/docs/go/getting-started

const connectGoServerTemplate = `

func main() {
	mux := http.NewServeMux()
	// The generated constructors return a path and a plain net/http
	// handler.
	%s
	err := http.ListenAndServe(
	  "localhost:8080",
	  // For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
	  // avoid x/net/http2 by using http.ListenAndServeTLS.
	  h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: " + err.Error())
  }
  
`

func GenConnectServer(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	services := []string{} // should probably be passed in

	for _, v := range file.Services {
		services = append(services, v.GoName) // service.Desc.Name()
	}

	fileName := strings.ToLower("cmd" + "/" + string(file.GoPackageName) + "/" + "main.go")
	g := gen.NewGeneratedFile(fileName, protogen.GoImportPath("."))

	g.P("package main ")

	// get connect, go import path

	// required imports
	//g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: file.GoImportPath})
	g.P("// " + protogen.GoIdent{GoImportPath: file.GoImportPath}.String())
	// example "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/example/exampleconnect" is what we want
	// example "github.com/lcmaguire/protoc-gen-go-goo/example" is what we get

	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "log"})
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "net/http"})
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "golang.org/x/net/http2"})
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "golang.org/x/net/http2/h2c"})

	// hard coding these vals for now, will have to think of a cleaner way of figuring out go mod path + out path for generated services.
	const _hardCodedPath = "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect" // + {{goName}} + gonameconnect

	goPKGname := strings.ToLower(string(file.GoPackageName))
	connectGenImportPath := fmt.Sprintf("%s/%s/%s", _hardCodedPath, goPKGname, goPKGname+"connect")
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: protogen.GoImportPath(connectGenImportPath)})

	pkg := getParamPKG(file.GoDescriptorIdent.GoImportPath.String())

	resgisteredServices := ""
	for _, serviceName := range services {
		// dir goModPath + serviceName
		// will also need to get go-goo_out path to put inbetween

		importPath := fmt.Sprintf("%s/%s", _hardCodedPath, strings.ToLower(serviceName))
		g.QualifiedGoIdent(protogen.GoIdent{GoName: "", GoImportPath: protogen.GoImportPath(importPath)})

		resgisteredServices += fmt.Sprintf(
			serviceHandleTemplate,
			pkg,
			serviceName,
			strings.ToLower(serviceName)+"."+serviceName,
		)
	}

	g.P(fmt.Sprintf(connectGoServerTemplate, resgisteredServices))

	return nil
}

// want {{protoPKG}}.New{{ServiceName}}Handler(&{{implemnentationAlias}}.{{ServerName}}{})

const serviceHandleTemplate = `

mux.Handle(%sconnect.New%sHandler(&%s{}))
`
