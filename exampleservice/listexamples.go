package exampleservice

import (
	context "context"
	out "github.com/lcmaguire/protoc-gen-go-goo/example/out"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func (e *ExampleService) ListExamples(ctx context.Context, in *out.SearchRequest) (out *out.SearchResponse, err error) {
	return nil, status.Error(codes.Unimplemented, "yet to be implemented")
}
