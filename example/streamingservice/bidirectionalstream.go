package streamingservice

import (
	"fmt"
	"io"
	"log"

	"github.com/lcmaguire/protoc-gen-go-goo/example"
)

// BiDirectionalStream implements tutorial.StreamingService.BiDirectionalStream. // TODO get example.StreamingService_BiDirectionalStreamServer from proto.
func (s *StreamingService) BiDirectionalStream(bidi example.StreamingService_BiDirectionalStreamServer) error {
	ctx := bidi.Context()
	for {

		// exit if context is done
		// or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		req, err := bidi.Recv()
		if err == io.EOF {
			// return will close stream from server side
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}
		fmt.Println(req)

		resp := &example.GreetResponse{}
		if err := bidi.Send(resp); err != nil {
			log.Printf("send error %v", err)
		}
	}

	// return status.Error(codes.Unimplemented, "yet to be implemented")
}
