package exampleservice

import (
	context "context"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	"google.golang.org/protobuf/types/known/emptypb"
	"github.com/lcmaguire/protoc-gen-go-goo/example"
)

// DeleteExample implements tutorial.ExampleService.DeleteExample.
func (e *ExampleService) DeleteExample(ctx context.Context, in *example.SearchRequest) (out *emptypb.Empty, err error) {
	return nil, status.Error(codes.Unimplemented, "yet to be implemented")
}
