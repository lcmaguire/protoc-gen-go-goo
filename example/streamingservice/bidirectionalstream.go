package streamingservice

import (
	context "context"
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// BiDirectionalStream implements tutorial.StreamingService.BiDirectionalStream.
func (s *StreamingService) BiDirectionalStream(stream example.StreamingService_BiDirectionalStreamServer) error {
	return status.Error(codes.Unimplemented, "yet to be implemented")
}
