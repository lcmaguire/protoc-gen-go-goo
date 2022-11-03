package exampleservice

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lcmaguire/protoc-gen-go-goo/example"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteExample implements tutorial.ExampleService.DeleteExample.
func (e *ExampleService) DeleteExample(ctx context.Context, in *example.SearchRequest) (out *emptypb.Empty, err error) {
	return nil, status.Error(codes.Unimplemented, "yet to be implemented")
}
