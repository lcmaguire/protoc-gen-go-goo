package exampleservice

import (
	context "context"
	out "example/out"
)

func (e *ExampleService) ListExamples(ctx context.Context, in *out.SearchRequest) (out *out.SearchResponse, err error) {
	return
}
