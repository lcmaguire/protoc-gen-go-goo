package exampleservice

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	example "github.com/lcmaguire/protoc-gen-go-goo/example"
)

// ListExamples implements tutorial.ExampleService.ListExamples.
func (e *ExampleService) ListExamples(ctx context.Context, in *example.SearchRequest) (out *example.SearchResponse, err error) {
	return nil, status.Error(codes.Unimplemented, "yet to be implemented")
}
