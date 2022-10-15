package templates

// Server stuff looks v different
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

const serviceHandleTemplate = `

mux.Handle(%sconnect.New%sHandler(&%s{}))
`

// METHOD

// can probably pass in  req  *connect_go.Request[%s] , *connect_go.Response[%s]

const methodTemplate = `

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
		res, err := service.%s(context.Background(), nil)
		assert.Error(t, err)
		assert.Equal(t, codes.Unimplemented, status.Code(err))
		assert.Nil(t, res)
	}
	`

// looks the same, can probably use default
const serviceTemplate = `
// %s ...
type %s struct { 
	%s.Unimplemented%sServer
}
	`

// can probably use defualt.
const methodCallerTemplate = `%s *%s`
