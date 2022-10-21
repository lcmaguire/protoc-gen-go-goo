package streamingservice

import (
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
)

// StreamingService implements tutorial.StreamingService.
type StreamingService struct {
	example.UnimplementedStreamingServiceServer
}

func NewStreamingService() *StreamingService {
	return &StreamingService{}
}
