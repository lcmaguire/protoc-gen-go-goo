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

func (%s) %s(ctx context.Context, req *connect_go.Request[%s]) (*connect_go.Response[%s], error) {
	res := connect_go.NewResponse(&%s{})
	return res, status.Error(codes.Unimplemented, "yet to be implemented")
}

`

// todo make this work
const TestFileTemplate = `
	func Test%s(t *testing.T){
		t.Parallel()
		service := &%s{}
		req := &connect_go.Request[%s]{
			Msg: &%s{},
		}
		res, err := service.%s(context.Background(), req)
		assert.Error(t, err)
		assert.Equal(t, codes.Unimplemented, status.Code(err))
		proto.Equal(res.Msg, &%s{})
	}
	`
