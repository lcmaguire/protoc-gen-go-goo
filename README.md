# protoc-gen-go-goo

a protoc plugin that will generate boilerplate go code for all gRPC services and methods defined in your proto file.

for all services included in your protoc command will generate a directory containing the following:
- File with struct that implements your service
- File for all RPC endpoints with all imports required for the service
- Test file for all RPC endpoints that asserts unimplemented


protoc-gen-go-goo currently Supports 

|                     | gRPC-go | connect-go |
|---------------------|---------|------------|
| unary RPC file      |    :white_check_mark:       |      :white_check_mark:      |
| unary RPC test      |   :white_check_mark:        |       :white_check_mark:       |
| streaming RPC files |         |       :white_check_mark:       |
| streaming RPC tests |         |      :white_check_mark:        |
| basic server gen    |     :white_check_mark:      |       :white_check_mark:       |


## TODO , mention 

- Streaming GEN currently CONNECT only.
- CONFIG explain
- update examples


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

## Params / Config

| Name | Type |   Use   | Default |
| ---- | --- |  --  | ------- |
|  tests  | bool |   to determine if you want test files + tests generated for your generated RPC methods   | true |
|  connectGo  | bool |   will use [connect-go](https://pkg.go.dev/github.com/bufbuild/connect-go) over normal grpc-go   | false |
|  server | bool     |  if true will generate a basic server that will run your rpc services (generatedPath should be set too)  | false |
|  generatedPath  | string | used by server to import the code generated by this plugin code. will be gomod path + goo_out e.g. generatedPath=github.com/lcmaguire/protoc-gen-go-goo/example| "" |

### Config Use

include all possible paths + links to examples.

## gRPC Server implmentation

todo