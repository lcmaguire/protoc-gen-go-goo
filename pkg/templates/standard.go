package templates

const ServiceTemplate = `
// {{.ServiceName}} implements {{.FullName}}.
type {{.ServiceName}} struct { 
	{{.Pkg}}.Unimplemented{{.ServiceName}}Server
}
	
func New{{.ServiceName}} () *{{.ServiceName}} {
	return &{{.ServiceName}}{}
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
	// {{.MethodName}} implements {{.FullName}}.
	func ({{.S1}}*{{.ServiceName}}) {{.MethodName}} (ctx context.Context, in *{{.RequestType}}) (out *{{.ResponseType}}, err error){
		return nil, status.Error(codes.Unimplemented, "yet to be implemented")
	}
`

const ClientStreamingMethod = `
	// {{.MethodName}} implements {{.FullName}}.
	func ({{.S1}}*{{.ServiceName}}) {{.MethodName}} ( ) error{
		return status.Error(codes.Unimplemented, "yet to be implemented")
	}
`

/*
func (s *StreamingService) ClientStream(stream example.StreamingService_ClientStreamServer) error {
	return status.Error(codes.Unimplemented, "yet to be implemented")
}
*/

const ServerStreamingMethod = `
	// {{.MethodName}} implements {{.FullName}}.
	func ({{.S1}}*{{.ServiceName}}) {{.MethodName}} (in *{{.RequestType}}, stream {{.ReqPkg}}.{{.ServiceName}}_{{.MethodName}}Server) error{

		  return status.Error(codes.Unimplemented, "yet to be implemented")
	}
`

//  {{ServiceName}}_{{MethodName}}Server

/*
func (s *StreamingService) ResponseStream(in *example.GreetRequest, stream example.StreamingService_ResponseStreamServer) error {
	return status.Error(codes.Unimplemented, "yet to be implemented")
}
*/

const BidirectionStreamingMethod = `
	// {{.MethodName}} implements {{.FullName}}.
	func ({{.S1}}*{{.ServiceName}}) {{.MethodName}} (stream example.StreamingService_BiDirectionalStreamServer) error {
		return status.Error(codes.Unimplemented, "yet to be implemented")
	}
`

/*

func (s *StreamingService) BiDirectionalStream(stream example.StreamingService_BiDirectionalStreamServer) error {
	return status.Error(codes.Unimplemented, "yet to be implemented")
}

*/

const MethodCallerTemplate = `%s *%s`
