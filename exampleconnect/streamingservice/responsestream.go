package streamingservice

import (
	"context"
	"errors"
	connect_go "github.com/bufbuild/connect-go"
	// sample "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"

	"github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
)

// ResponseStream implements tutorial.StreamingService.ResponseStream.
func (s *StreamingService) ResponseStream(ctx context.Context, req *connect_go.Request[sample.GreetRequest]) (*connect_go.Response[sample.GreetResponse], error) {
	res := connect_go.NewResponse(&sample.GreetResponse{})
	return res, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("not yet implemented"))
}
