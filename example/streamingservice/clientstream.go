package streamingservice

import (
	context "context"
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// ClientStream implements tutorial.StreamingService.ClientStream.
func (s *StreamingService) ClientStream() error {
	return status.Error(codes.Unimplemented, "yet to be implemented")
}
