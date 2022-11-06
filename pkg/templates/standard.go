package templates

const ServiceTemplate = `

package {{.GoPkgName}}

import (
	{{.Imports}}
)

// {{.ServiceName}} implements {{.FullName}}.
type {{.ServiceName}} struct { 
	{{.Pkg}}.Unimplemented{{.ServiceName}}Server
}
	
func New{{.ServiceName}} () *{{.ServiceName}} {
	return &{{.ServiceName}}{}
}
`

// ServerTemplate template for a gRPC server that runs your service.
const ServerTemplate = `
package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"


	"{{.GenImportPath}}"

	{{.ServiceImports}}
)

func main() {
    if err := run(); err != nil {
        log.Fatal(err)
    }
}

func run() error {
    listenOn := "127.0.0.1:8080" // this should be passed in via config
    listener, err := net.Listen("tcp", listenOn) 
    if err != nil {
        return  err 
    }

    server := grpc.NewServer()
	// services in your protoFile
    {{.Services}}
	reflection.Register(server) // this should perhaps be optional
	log.Println("Listening on", listenOn)
    if err := server.Serve(listener); err != nil {
        return err 
    }

    return nil
}

`

// RegisterServiceTemplate handles registering your service + reflection for that endpoint.
const RegisterServiceTemplate = `
{{.Pkg}}.Register{{.ServiceName}}Server(server, &{{.ServiceStruct}}{})
`

const TestFileTemplate = `
package {{.GoPkgName}}

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"

	{{.Imports}}
)

func Test{{.MethodName}}(t *testing.T){
	t.Parallel()
	service := &{{.ServiceName}}{}
	req := &{{.RequestType}}{}
	res, err := service.{{.MethodName}}(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, codes.Unimplemented, status.Code(err))
	proto.Equal(res, &{{.ResponseType}}{})
}
`

const MethodTemplate = `
package {{.GoPkgName}}

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	{{.Imports}}
)

// {{.MethodName}} implements {{.FullName}}.
func ({{.MethodCaller}}) {{.MethodName}} (ctx context.Context, in *{{.RequestType}}) (out *{{.ResponseType}}, err error){
	return nil, status.Error(codes.Unimplemented, "yet to be implemented")
}
`

const MethodCallerTemplate = `%s *%s`
