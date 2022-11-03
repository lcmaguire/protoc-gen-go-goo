package exampleservice

import (
	context "context"
	connect_go "github.com/bufbuild/connect-go"
	proto "github.com/golang/protobuf/proto"
	assert "github.com/stretchr/testify/assert"
	testing "testing"

	"github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestDeleteExample(t *testing.T) {
	t.Parallel()
	service := &ExampleService{}
	req := &connect_go.Request[sample.SearchRequest]{
		Msg: &sample.SearchRequest{},
	}
	res, err := service.DeleteExample(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, connect_go.CodeUnimplemented, connect_go.CodeOf(err))
	proto.Equal(res.Msg, &emptypb.Empty{})
}
