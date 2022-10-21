package streamingservice

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	sample "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
)

// ResponseStream implements tutorial.StreamingService.ResponseStream.
func (s *StreamingService) ResponseStream(ctx context.Context, req *connect_go.Request[sample.GreetRequest], stream *connect_go.ServerStream[sample.GreetResponse]) error {
	/* ticker := time.NewTicker(time.Second) // You should set this via config.
	defer ticker.Stop()
	for i := 0; i &lt; 5; i++ {
		if ticker != nil {
			select {
			case &lt;-ctx.Done():
				return ctx.Err()
			case &lt;-ticker.C:
			}
		}
		if err := stream.Send(&sample.GreetResponse{}); err != nil {
			return err
		}
	}*/
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("not yet implemented"))
}
