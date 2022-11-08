package streamingservice

import (
	"io"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lcmaguire/protoc-gen-go-goo/example"
)

// ClientStream implements tutorial.StreamingService.ClientStream. // todo get clientStream pkg.ServiceName_MethodNameServer?
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
