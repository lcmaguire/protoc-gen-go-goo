package exampleservice

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	example "github.com/lcmaguire/protoc-gen-go-goo/example"
)

// CreateExample implements tutorial.ExampleService.CreateExample.
func (e *ExampleService) CreateExample(ctx context.Context, in *example.SearchRequest) (out *example.SearchResponse, err error) {
	return nil, status.Error(codes.Unimplemented, "yet to be implemented")
}
