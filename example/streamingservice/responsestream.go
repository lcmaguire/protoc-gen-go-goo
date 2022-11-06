package streamingservice

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lcmaguire/protoc-gen-go-goo/example"
)

// ResponseStream implements tutorial.StreamingService.ResponseStream.
func (s *StreamingService) ResponseStream(ctx context.Context, in *example.GreetRequest) (out *example.GreetResponse, err error) {
	return nil, status.Error(codes.Unimplemented, "yet to be implemented")
}
