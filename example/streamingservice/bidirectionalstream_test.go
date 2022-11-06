package streamingservice

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"

	"github.com/lcmaguire/protoc-gen-go-goo/example"
	"github.com/lcmaguire/protoc-gen-go-goo/exampleconnect"
)

func TestBiDirectionalStream(t *testing.T) {
	t.Parallel()
	service := &StreamingService{}
	req := &example.GreetRequest{}
	res, err := service.BiDirectionalStream(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, codes.Unimplemented, status.Code(err))
	proto.Equal(res, &example.GreetResponse{})
}
