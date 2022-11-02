package streamingservice

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"

	"github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
)

// ClientStream implements tutorial.StreamingService.ClientStream.
func (s *StreamingService) ClientStream(ctx context.Context, stream *connect_go.ClientStream[sample.GreetRequest]) (*connect_go.Response[sample.GreetResponse], error) {
	for stream.Receive() {
		// implement logic here.
	}
	if err := stream.Err(); err != nil {
		return nil, connect_go.NewError(connect_go.CodeUnknown, err)
	}
	res := connect_go.NewResponse(&sample.GreetResponse{})
	return res, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("not yet implemented"))
}
