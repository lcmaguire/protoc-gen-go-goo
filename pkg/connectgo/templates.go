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
func ({{.S1}}*{{.ServiceName}}) {{.MethodName}}(ctx context.Context, req *connect_go.Request[{{.RequestType}}]) (*connect_go.Response[{{.ResponseType}}], error) {
	res := connect_go.NewResponse(&{{.ResponseType}}{})
	return res, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("not yet implemented"))
}

`

const ClientStreamingTemplate = `
// {{.MethodName}} implements {{.FullName}}.
func ({{.S1}}*{{.ServiceName}}) {{.MethodName}}(ctx context.Context, stream *connect_go.ClientStream[{{.RequestType}}]) (*connect_go.Response[{{.ResponseType}}], error) {
	for stream.Receive() {
		// implement logic here.
	}
	if err := stream.Err(); err != nil {
	  return nil, connect_go.NewError(connect_go.CodeUnknown, err)
	}
	res := connect_go.NewResponse(&{{.ResponseType}}{})
	return res, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("not yet implemented")) 
  }
`

const ServerStreamingTemplate = `
// {{.MethodName}} implements {{.FullName}}.
func ({{.S1}}*{{.ServiceName}}) {{.MethodName}}(ctx context.Context, req *connect_go.Request[{{.RequestType}}], stream *connect_go.ServerStream[{{.ResponseType}}]) error {
	ticker := time.NewTicker(time.Second) // You should set this via config.
	defer ticker.Stop()
	for i := 0; i < 5 ; i++ {
		if ticker != nil {
			select {
			case <- ctx.Done():
				return ctx.Err()
			case <- ticker.C:
			}
		}
		if err := stream.Send(&{{.ResponseType}}{}); err != nil {
			return err
		}
	}
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("not yet implemented"))
}
`

const BiDirectionalStreamingTemplate = `
// {{.MethodName}} implements {{.FullName}}.
func ({{.S1}}*{{.ServiceName}}) {{.MethodName}}(ctx context.Context, stream *connect_go.BidiStream[{{.RequestType}}, {{.ResponseType}}]) error {
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		request, err := stream.Receive()
		if err != nil && errors.Is(err, io.EOF) {
			return nil
		} else if err != nil {
			return err
		}
		fmt.Println("incoming request ", request)
		if err := stream.Send(&{{.ResponseType}}{}); err != nil {
			return err
		}
		connect_go.NewError(connect_go.CodeUnimplemented, errors.New("not yet implemented"))
	}
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
		assert.Equal(t, connect_go.CodeUnimplemented, connect_go.CodeOf(err))
		proto.Equal(res.Msg, &{{.ResponseType}}{})
	}
	`

// %sconnect.New%sHandler(&%s{})
const TestBidirectionalStreamTemplate = `
func Test{{.MethodName}}(t *testing.T){
	t.Parallel()
	mux := http.NewServeMux()

	mux.Handle({{.Pkg}}connect.New{{.ServiceName}}Handler(&{{.ServiceName}}{}))
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	defer server.Close()

	connectClient := {{.Pkg}}connect.New{{.ServiceName}}Client(
		server.Client(),
		server.URL,
	)
	grpcClient := {{.Pkg}}connect.New{{.ServiceName}}Client(
		server.Client(),
		server.URL,
		connect_go.WithGRPC(),
	)
	clients := []{{.Pkg}}connect.{{.ServiceName}}Client{connectClient, grpcClient}

	t.Run("bidirectionalTest", func(t *testing.T) {
		for _, client := range clients {
			sendValues := []string{"Hello!", "How are you doing?", "I have an issue with my bike", "bye"}
			var receivedValues []string
			stream := client.{{.MethodName}}(context.Background())
			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				defer wg.Done()
				for _, sentence := range sendValues {
					err := stream.Send(&{{.RequestType}}{})
					require.Nil(t, err )
					fmt.Println(sentence)
				}
				require.Nil(t, stream.CloseRequest())
			}()
			go func() {
				defer wg.Done()
				for {
					_, err := stream.Receive()
					if errors.Is(err, io.EOF) {
						break
					}
					require.Nil(t, err)
					receivedValues = append(receivedValues, "")
				}
				require.Nil(t, stream.CloseResponse())
			}()
			wg.Wait()
			assert.Equal(t, len(receivedValues), len(sendValues))
		}
	})
}
	`

const TestClientStreamTemplate = `
func TestClientStream(t *testing.T) {	
	t.Parallel()
	mux := http.NewServeMux()

	mux.Handle({{.Pkg}}connect.New{{.ServiceName}}Handler(&{{.ServiceName}}{}))
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	defer server.Close()

	connectClient := {{.Pkg}}connect.New{{.ServiceName}}Client(
		server.Client(),
		server.URL,
	)
	grpcClient := {{.Pkg}}connect.New{{.ServiceName}}Client(
		server.Client(),
		server.URL,
		connect_go.WithGRPC(),
	)
	clients := []{{.Pkg}}connect.{{.ServiceName}}Client{connectClient, grpcClient}

	for _, client := range clients {
		sendValues := []string{"Hello!", "How are you doing?", "I have an issue with my bike", "bye"}
		stream := client.{{.MethodName}}(context.Background())
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, sentence := range sendValues {
				err := stream.Send(&{{.RequestType}}{})
				
				require.Nil(t, err, sentence)
			}
		}()
		wg.Wait()
		res, err := stream.CloseAndReceive()
		require.Error(t, err)
		assert.Equal(t, connect_go.CodeUnimplemented, connect_go.CodeOf(err))
		assert.Nil(t, res)
	}
}
	`

const UnsportedTestFile = `
func Test{{.MethodName}}(t *testing.T){
	t.Parallel()
	// tests for this type of RPC yet to be implemented.
	assert.NotNil(t, &{{.RequestType}}{})
	assert.NotNil(t, &{{.ResponseType}}{})
}
`
const ServiceTemplate = `
	// {{.ServiceName}} implements {{.FullName}}.
	type {{.ServiceName}} struct { 
		{{.Pkg}}.Unimplemented{{.ServiceName}}Handler
	}
		
	func New{{.ServiceName}} () *{{.ServiceName}} {
		return &{{.ServiceName}}{}
	}
	`
