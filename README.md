# protoc-gen-go-goo

a protoc plugin that will generate boilerplate go code for all gRPC services and methods defined in your proto file.

for all services included in your protoc command will generate a directory containing the following:
- File with struct that implements your service
- File for all RPC endpoints with all imports required for the service
- Test file for all RPC endpoints that asserts unimplemented


protoc-gen-go-goo currently Supports 

|                     | gRPC-go | connect-go |
|---------------------|---------|------------|
| unary RPC files      |    :white_check_mark:       |      :white_check_mark:      |
| unary RPC tests      |   :white_check_mark:        |       :white_check_mark:       |
| streaming RPC files |         |       :white_check_mark:       |
| streaming RPC tests |         |       :white_check_mark:        |
| basic server gen    |     :white_check_mark:      |       :white_check_mark:       |

### Example with gRPC-go


```
# generate only goo generated code, can also include 
    protoc -I=example \
    --go-goo_out=example \

    example/*.proto 
```

Example Output

A struct that implements your service.
```
// example/exampleservice/exampleservice.go
package exampleservice

import (
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
)

// ExampleService implements tutorial.ExampleService.
type ExampleService struct {
	example.UnimplementedExampleServiceServer
}

func NewExampleService() *ExampleService {
	return &ExampleService{}
}
```

files for all RPC methods for your service (below is an example of just one.) see a full example [here](./example).

```
// example/exampleservice/createexample.go
package exampleservice

import (
	context "context"
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// CreateExample implements tutorial.ExampleService.CreateExample.
func (e *ExampleService) CreateExample(ctx context.Context, in *example.SearchRequest) (out *example.SearchResponse, err error) {
	return nil, status.Error(codes.Unimplemented, "yet to be implemented")
}
```

and a test for your RPC method

```
// example/exampleservice/createexample_test.go
package exampleservice

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
	assert "github.com/stretchr/testify/assert"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	testing "testing"
)

func TestDeleteExample(t *testing.T) {
	t.Parallel()
	service := &ExampleService{}
	req := &example.SearchRequest{}
	res, err := service.DeleteExample(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, codes.Unimplemented, status.Code(err))
	proto.Equal(res, &emptypb.Empty{})
}
```

### Example Connect Go

see full example [here](./exampleconnect/)

```
# generate only goo generated code, can also include 
    protoc -I=exampleconnect \
	--go-goo_out=exampleconnect \
	--go-goo_opt=connectGo=true, \
	exampleconnect/example.proto 
```

Example Output

A struct that implements your service.
```
// exampleconnect/exampleservice/exampleservice.go
package exampleservice

import (
	sampleconnect "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sampleconnect"
)

// ExampleService implements tutorial.ExampleService.
type ExampleService struct {
	sampleconnect.UnimplementedExampleServiceHandler
}

func NewExampleService() *ExampleService {
	return &ExampleService{}
}
```

files for all RPC methods for your service (below is an example of just one.) see a full example [here](./exampleconnect).

```
// exampleconnect/exampleservice/createexample.go
package exampleservice

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	sample "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
)

// CreateExample implements tutorial.ExampleService.CreateExample.
func (e *ExampleService) CreateExample(ctx context.Context, req *connect_go.Request[sample.SearchRequest]) (*connect_go.Response[sample.SearchResponse], error) {
	res := connect_go.NewResponse(&sample.SearchResponse{})
	return res, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("not yet implemented"))
}
```

and a test for your RPC method

```
// exampleconnect/exampleservice/createexample_test.go
package exampleservice

import (
	context "context"
	connect_go "github.com/bufbuild/connect-go"
	proto "github.com/golang/protobuf/proto"
	sample "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
	assert "github.com/stretchr/testify/assert"
	testing "testing"
)

func TestCreateExample(t *testing.T) {
	t.Parallel()
	service := &ExampleService{}
	req := &connect_go.Request[sample.SearchRequest]{
		Msg: &sample.SearchRequest{},
	}
	res, err := service.CreateExample(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, connect_go.CodeUnimplemented, connect_go.CodeOf(err))
	proto.Equal(res.Msg, &sample.SearchResponse{})
}
```

for connect go it also supports generating boiler plate code for streaming RPC endpoints you can view example generated code [here](./exampleconnect/streamingservice/).

below is an example generated BiDirectional streaming RPC (also will gen boiler plate for ClientStreaming and ServerStreaming RPC's)

```
package streamingservice

import (
	context "context"
	errors "errors"
	fmt "fmt"
	connect_go "github.com/bufbuild/connect-go"
	sample "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
	io "io"
)

// BiDirectionalStream implements tutorial.StreamingService.BiDirectionalStream.
func (s *StreamingService) BiDirectionalStream(ctx context.Context, stream *connect_go.BidiStream[sample.GreetRequest, sample.GreetResponse]) error {
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
		if err := stream.Send(&sample.GreetResponse{}); err != nil {
			return err
		}
		connect_go.NewError(connect_go.CodeUnimplemented, errors.New("not yet implemented"))
	}
}
```

and its test file 

```
package streamingservice

import (
	context "context"
	errors "errors"
	fmt "fmt"
	connect_go "github.com/bufbuild/connect-go"
	sample "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
	sampleconnect "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sampleconnect"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	io "io"
	http "net/http"
	httptest "net/http/httptest"
	sync "sync"
	testing "testing"
)

func TestBiDirectionalStream(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()

	mux.Handle(sampleconnect.NewStreamingServiceHandler(&StreamingService{}))
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	defer server.Close()

	connectClient := sampleconnect.NewStreamingServiceClient(
		server.Client(),
		server.URL,
	)
	grpcClient := sampleconnect.NewStreamingServiceClient(
		server.Client(),
		server.URL,
		connect_go.WithGRPC(),
	)
	clients := []sampleconnect.StreamingServiceClient{connectClient, grpcClient}

	t.Run("bidirectionalTest", func(t *testing.T) {
		for _, client := range clients {
			sendValues := []string{"Hello!", "How are you doing?", "I have an issue with my bike", "bye"}
			var receivedValues []string
			stream := client.BiDirectionalStream(context.Background())
			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				defer wg.Done()
				for _, sentence := range sendValues {
					err := stream.Send(&sample.GreetRequest{})
					require.Nil(t, err)
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
```


## Options

| Name | Type |   Use   | Default |
| ---- | --- |  --  | ------- |
|  tests  | bool |   to determine if you want test files + tests generated for your generated RPC methods   | true |
|  connectGo  | bool |   will use [connect-go](https://pkg.go.dev/github.com/bufbuild/connect-go) over normal grpc-go   | false |
|  server | bool     |  if true will generate a basic server that will run your rpc services (generatedPath should be set too)  | false |
|  generatedPath  | string | used by server to import the code generated by this plugin code. will be gomod path + goo_out e.g. generatedPath=github.com/lcmaguire/protoc-gen-go-goo/example| "" |

## gRPC Server implmentation

To also generate a basic gRPC server you can pass in the following options.

```
# this pattern
--go-goo_opt=server=true,generatedPath={{out path go module}}/{{go_goo-out}}

# go-gRPC example
--go-goo_opt=server=true,generatedPath=github.com/lcmaguire/protoc-gen-go-goo/example

# connect-go example
--go-goo_opt=server=true,connectGo=true,generatedPath=github.com/lcmaguire/protoc-gen-go-goo/exampleconnect

# full command
protoc -I=exampleconnect \
	--go-goo_out=exampleconnect \
	--go-goo_opt=tests=true,server=true,connectGo=true,generatedPath=github.com/lcmaguire/protoc-gen-go-goo/exampleconnect \
	exampleconnect/example.proto 
	

```
it will generate a sample server like below


[example for go-grpc](./example/cmd/example/main.go)

```
package main

import (
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
	exampleservice "github.com/lcmaguire/protoc-gen-go-goo/example/exampleservice"
	grpc "google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection"
	log "log"
	net "net"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := "127.0.0.1:8080"                 // this should be passed in via config
	listener, err := net.Listen("tcp", listenOn) // this too
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	// services in your protoFile

	example.RegisterExampleServiceServer(server, &exampleservice.ExampleService{})
	reflection.Register(server) // this should perhaps be optional

	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return err
	}

	return nil
}
```

[example for connect-go](./exampleconnect/cmd/sample/main.go)

```
package main

import (
	exampleservice "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/exampleservice"
	sampleconnect "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sampleconnect"
	streamingservice "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/streamingservice"
	http2 "golang.org/x/net/http2"
	h2c "golang.org/x/net/http2/h2c"
	log "log"
	http "net/http"
)

func main() {
	mux := http.NewServeMux()
	// The generated constructors return a path and a plain net/http
	// handler.

	mux.Handle(sampleconnect.NewExampleServiceHandler(&exampleservice.ExampleService{}))

	mux.Handle(sampleconnect.NewStreamingServiceHandler(&streamingservice.StreamingService{}))

	err := http.ListenAndServe(
		"localhost:8080",
		// For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
		// avoid x/net/http2 by using http.ListenAndServeTLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: " + err.Error())
}
```
