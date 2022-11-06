package streamingservice

import (
	"io"

	"github.com/lcmaguire/protoc-gen-go-goo/example"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ClientStream implements tutorial.StreamingService.ClientStream.
func (s *StreamingService) ClientStream(stream example.StreamingService_ClientStreamServer) error {
	// ctx := stream.Context()
	// var req *example.GreetRequest
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return status.Error(codes.Unimplemented, stream.SendAndClose(&example.GreetResponse{}).Error())
		}
		if err != nil {
			return err
		}
	}
}
