package streamingservice

import (
	sampleconnect "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sampleconnect"
)

// StreamingService implements tutorial.StreamingService.
type StreamingService struct {
	sampleconnect.UnimplementedStreamingServiceHandler
}

func NewStreamingService() *StreamingService {
	return &StreamingService{}
}
