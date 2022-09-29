package exampleservice

import (
	context "context"
	out "github.com/lcmaguire/protoc-gen-go-goo/github.com/lcmaguire/protoc-gen-go-goo/out"
)

func (e *ExampleService) GetExample(ctx context.Context, in *out.SearchRequest) (out *out.SearchResponse, err error) {
	return
}
