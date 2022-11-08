package streamingservice

import (
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lcmaguire/protoc-gen-go-goo/example"
)

// ResponseStream implements tutorial.StreamingService.ResponseStream. // todo get ResponseStreamServer pkg.ServiceName_MethodNameServer?
func (s *StreamingService) ResponseStream(req *example.GreetRequest, stream example.StreamingService_ResponseStreamServer) error {
	ctx := stream.Context()
	ticker := time.NewTicker(time.Second) // You should set this via config.
	defer ticker.Stop()
	for i := 0; i < 5; i++ {
		if ticker != nil {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-ticker.C:
			}
		}
		if err := stream.Send(&example.GreetResponse{}); err != nil {
			return err
		}
	}
	return status.Error(codes.Unimplemented, "yet to be implemented.")
}
