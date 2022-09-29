package exampleservice

import (
	context "context"
	out "example/out"
)

func (e *ExampleService) GetExample(ctx context.Context, in *out.SearchRequest) (out *out.SearchResponse, err error) {
	return
}
