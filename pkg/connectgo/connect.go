package connectgo

// service appears to be the same

// rpc methods recieve/return connect.Request, connect.NewResponse

// server is much simpler

// import connectgo

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// sampled from https://connect.build/docs/go/getting-started

const connectGoServerTemplate = `

package main

import (
  "context"
  "log"
  "net/http"

  "github.com/bufbuild/connect-go"
  "golang.org/x/net/http2"
  "golang.org/x/net/http2/h2c"
)

func main() {
	mux := http.NewServeMux()
	// The generated constructors return a path and a plain net/http
	// handler.
	// want {{PKG}}.New{{ServiceName}}Handler(&{{importAlias}}.{{ServerName}}{})
	mux.Handle(pingv1connect.NewPingServiceHandler(&PingServer{}))
	err := http.ListenAndServe(
	  "localhost:8080",
	  // For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
	  // avoid x/net/http2 by using http.ListenAndServeTLS.
	  h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: " + err)
  }
  
`

func GenConnectServer() *protogen.GeneratedFile {

	return nil
}

// takes in ServiceName, Request, response,
const connectGoRPCMethodTemplate = `

func (%s) Ping(ctx context.Context, req *connect.Request[%s]) (*connect.Response[%s], error) {
	res := connect.NewResponse(&pingv1.PingResponse{})
	return res, nil
}

`

// pass in
func GenConnectRPCMethod() *protogen.GeneratedFile {
	return nil
}

func GenerateFilesForService(gen *protogen.Plugin, service *protogen.Service, file *protogen.File) *protogen.GeneratedFile {
	// need to abstract stuff away, will do later
	generateServiceFile(gen, service)

	for _, v := range service.Methods {

		genRpcMethod(gen, service, v)
		//outfiles = append(outfiles, g)

		// wil generate test file
		//genTestFile(gen, service, v)
		//outfiles = append(outfiles, gT)
	}

	return nil
}

func ConnectGen(gen *protogen.Plugin) {
	//
	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}

		for _, v := range f.Services {
			GenerateFilesForService(gen, v, f)
		}

		GenConnectServer()
	}

}
