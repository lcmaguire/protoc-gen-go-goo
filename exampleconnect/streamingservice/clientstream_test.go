package streamingservice

import (
	context "context"
	connect_go "github.com/bufbuild/connect-go"
	proto "github.com/golang/protobuf/proto"
	sample "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
	assert "github.com/stretchr/testify/assert"
	testing "testing"
)

func TestClientStream(t *testing.T) {
	t.Parallel()
	service := &StreamingService{}
	req := &connect_go.Request[sample.GreetRequest]{
		Msg: &sample.GreetRequest{},
	}
	res, err := service.ClientStream(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, connect_go.CodeUnimplemented, connect_go.CodeOf(err))
	proto.Equal(res.Msg, &sample.GreetResponse{})
}
