package exampleservice

import (
	context "context"
	connect_go "github.com/bufbuild/connect-go"
	sample "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// DeleteExamples implements tutorial.ExampleService.DeleteExamples.
func (e *ExampleService) DeleteExamples(ctx context.Context, req *connect_go.Request[sample.SearchRequest]) (*connect_go.Response[sample.SearchResponse], error) {
	res := connect_go.NewResponse(&sample.SearchResponse{})
	return res, status.Error(codes.Unimplemented, "yet to be implemented")
}
