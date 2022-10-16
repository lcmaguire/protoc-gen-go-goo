package connectgo

// Server stuff looks v different
// sampled from https://connect.build/docs/go/getting-started

// ConnectGoServerTemplate ...
const ConnectGoServerTemplate = `

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

// ServiceHandleTemplate ...
const ServiceHandleTemplate = `

mux.Handle(%sconnect.New%sHandler(&%s{}))
`

// MethodTemplate ...
const MethodTemplate = `
// {{.MethodName}} implements {{.FullName}}.
func ({{.MethodCaller}}) {{.MethodName}}(ctx context.Context, req *connect_go.Request[{{.RequestType}}]) (*connect_go.Response[{{.ResponseType}}], error) {
	res := connect_go.NewResponse(&{{.ResponseType}}{})
	return res, status.Error(codes.Unimplemented, "yet to be implemented")
}

`

const TestFileTemplate = `
	func Test{{.MethodName}}(t *testing.T){
		t.Parallel()
		service := &{{.ServiceName}}{}
		req := &connect_go.Request[{{.RequestType}}]{
			Msg: &{{.RequestType}}{},
		}
		res, err := service.{{.MethodName}}(context.Background(), req)
		assert.Error(t, err)
		assert.Equal(t, codes.Unimplemented, status.Code(err))
		proto.Equal(res.Msg, &{{.ResponseType}}{})
	}
	`
