package streamingservice

import (
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// ClientStream implements tutorial.StreamingService.ClientStream.
func (s *StreamingService) ClientStream(example.StreamingService_ClientStreamServer) error {
	return status.Error(codes.Unimplemented, "yet to be implemented")
}
