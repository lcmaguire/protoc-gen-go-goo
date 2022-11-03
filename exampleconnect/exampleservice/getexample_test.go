package exampleservice

import (
	"context"
	connect_go "github.com/bufbuild/connect-go"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
)

func TestGetExample(t *testing.T) {
	t.Parallel()
	service := &ExampleService{}
	req := &connect_go.Request[sample.SearchRequest]{
		Msg: &sample.SearchRequest{},
	}
	res, err := service.GetExample(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, connect_go.CodeUnimplemented, connect_go.CodeOf(err))
	proto.Equal(res.Msg, &sample.SearchResponse{})
}
