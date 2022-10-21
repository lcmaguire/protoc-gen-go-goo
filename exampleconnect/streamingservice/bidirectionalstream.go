package streamingservice

import (
	context "context"
	errors "errors"
	fmt "fmt"
	connect_go "github.com/bufbuild/connect-go"
	sample "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
	io "io"
)

// BiDirectionalStream implements tutorial.StreamingService.BiDirectionalStream.
func (s *StreamingService) BiDirectionalStream(ctx context.Context, stream *connect_go.BidiStream[sample.GreetRequest, sample.GreetResponse]) error {
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		request, err := stream.Receive()
		if err != nil && errors.Is(err, io.EOF) {
			return nil
		} else if err != nil {
			return err
		}
		fmt.Println("incoming request ", request)
		if err := stream.Send(&sample.GreetResponse{}); err != nil {
			return err
		}
		connect_go.NewError(connect_go.CodeUnimplemented, errors.New("not yet implemented"))
	}
}
