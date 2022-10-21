package streamingservice

import (
	context "context"
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// ResponseStream implements tutorial.StreamingService.ResponseStream.
func (s *StreamingService) ResponseStream(in *example.GreetRequest, stream GreetRequest.StreamingService_ResponseStreamServer) error {

	return status.Error(codes.Unimplemented, "yet to be implemented")
}
