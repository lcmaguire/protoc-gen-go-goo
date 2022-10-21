package exampleservice

import (
	context "context"
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// DeleteExamples implements tutorial.ExampleService.DeleteExamples.
func (e *ExampleService) DeleteExamples(ctx context.Context, in *example.SearchRequest) (out *emptypb.Empty, err error) {
	return nil, status.Error(codes.Unimplemented, "yet to be implemented")
}
