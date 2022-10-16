package templates

const ServiceTemplate = `
// %s ...
type %s struct { 
	%s.Unimplemented%sServer
}
	`

const ActualServiceTemplate = `
// {{.ServiceName}} ...
type {{.ServiceName}} struct { 
	{{.Pkg}}.Unimplemented{{.ServiceName}}Server
}
	
`

// add in reflection api
const ServerTemplate = `
func main() {
    if err := run(); err != nil {
        log.Fatal(err)
    }
}

func run() error {
    listenOn := "127.0.0.1:8080" // this should be passed in via config
    listener, err := net.Listen("tcp", listenOn) // this too
    if err != nil {
        return  err 
    }

    server := grpc.NewServer()
	// services in your protoFile
    %s
	log.Println("Listening on", listenOn)
    if err := server.Serve(listener); err != nil {
        return err 
    }

    return nil
}

`

const RegisterServiceTemplate = `
%s.Register%sServer(server, &%s{})
reflection.Register(server) // this should perhaps be optional

`

const TestFileTemplate = `
	func Test%s(t *testing.T){
		t.Parallel()
		service := &%s{}
		req := &%s{}
		res, err := service.%s(context.Background(), req)
		assert.Error(t, err)
		assert.Equal(t, codes.Unimplemented, status.Code(err))
		proto.Equal(res, &%s{})
	}
	`

const MethodTemplate = `
	// %s ...
	func (%s) %s (ctx context.Context, in *%s) (out *%s, err error){
		return nil, status.Error(codes.Unimplemented, "yet to be implemented")
	}
`

const MethodCallerTemplate = `%s *%s`
