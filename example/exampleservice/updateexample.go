package exampleservice

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	example "github.com/lcmaguire/protoc-gen-go-goo/example"
)

// UpdateExample implements tutorial.ExampleService.UpdateExample.
func (s *ExampleService) UpdateExample(ctx context.Context, in *example.SearchRequest) (out *example.SearchResponse, err error) {
	return nil, status.Error(codes.Unimplemented, "yet to be implemented")
}
