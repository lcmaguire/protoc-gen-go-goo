package exampleservice

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	example "github.com/lcmaguire/protoc-gen-go-goo/example"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// DeleteExample implements tutorial.ExampleService.DeleteExample.
func (s *ExampleService) DeleteExample(ctx context.Context, in *example.SearchRequest) (out *emptypb.Empty, err error) {
	return nil, status.Error(codes.Unimplemented, "yet to be implemented")
}
