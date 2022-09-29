package exampleservice

import (
	context "context"
	out "github.com/lcmaguire/protoc-gen-go-goo/example/out"
)

func (e *ExampleService) ListExamples(ctx context.Context, in *out.SearchRequest) (out *out.SearchResponse, err error) {
	return
}
