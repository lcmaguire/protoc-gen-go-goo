package exampleservice

import (
	context "context"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	"github.com/lcmaguire/protoc-gen-go-goo/example"
)

// UpdateExample implements tutorial.ExampleService.UpdateExample.
func (e *ExampleService) UpdateExample(ctx context.Context, in *example.SearchRequest) (out *example.SearchResponse, err error) {
	return nil, status.Error(codes.Unimplemented, "yet to be implemented")
}
