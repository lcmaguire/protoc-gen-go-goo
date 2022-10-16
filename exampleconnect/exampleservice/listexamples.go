package exampleservice

import (
	context "context"
	connect_go "github.com/bufbuild/connect-go"
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// ListExamples implements tutorial.ExampleService.ListExamples.
func (e *ExampleService) ListExamples(ctx context.Context, req *connect_go.Request[example.SearchRequest]) (*connect_go.Response[example.SearchResponse], error) {
	res := connect_go.NewResponse(&example.SearchResponse{})
	return res, status.Error(codes.Unimplemented, "yet to be implemented")
}
