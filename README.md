# protoc-gen-go-goo

a protoc plugin that will generate boilerplate go code for all gRPC services and methods defined in your proto file.

for all services included in your protoc command will generate a directory containing the following:
- File with struct that implements your service
- File for all RPC endpoints with all imports required for the service
- Test file for all RPC endpoints that asserts unimplemented



### Example Command


```
# generate only goo generated code, can also include --go-goo_opt=
go install . && \
    protoc -I=example \
    --go-goo_out=. \
    *.proto 

```

Example Output

A struct that implements your service.
```
// example/exampleservice/exampleservice.go
package exampleservice

import (
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
)

// ExampleService ...
type ExampleService struct {
	example.UnimplementedExampleServiceServer
}

```

files for all RPC methods for your service

```
// example/exampleservice/createexample.go
package exampleservice

import (
	context "context"
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// CreateExample ...
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
	assert "github.com/stretchr/testify/assert"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	testing "testing"
)

func TestCreateExample(t *testing.T) {
	t.Parallel()
	service := &ExampleService{}
	res, err := service.CreateExample(context.Background(), nil)
	assert.Error(t, err)
	assert.Equal(t, codes.Unimplemented, status.Code(err))
	assert.Nil(t, res)
}

```

## Params / Config

| Name | Type |   Use   | Default |
| ---- | --- |  --  | ------- |
|  tests  | bool |   to determine if you want test files + tests generated for your generated RPC methods   | true |
|  server | bool     |  if true will generate a basic server that will run your rpc services  | true |
|  connectGo  | bool |   will use [connect-go](https://pkg.go.dev/github.com/bufbuild/connect-go) over normal grpc-go   | false |


## gRPC Server implmentation

todo