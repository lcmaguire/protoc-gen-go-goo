package exampleservice

import (
	context "context"
	"fmt"
	testing "testing"

	connect_go "github.com/bufbuild/connect-go"
	proto "github.com/golang/protobuf/proto"
	sample "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
	assert "github.com/stretchr/testify/assert"
)

func TestCreateExample(t *testing.T) {
	t.Parallel()
	service := &ExampleService{}
	req := &connect_go.Request[sample.SearchRequest]{
		Msg: &sample.SearchRequest{},
	}
	res, err := service.CreateExample(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, connect_go.CodeUnimplemented, connect_go.CodeOf(err))
	proto.Equal(res.Msg, &sample.SearchResponse{})

	fmt.Println(sample.File_example_proto.Options().ProtoReflect().Descriptor().FullName())
	panic(sample.File_example_proto.Services().Get(0).FullName())
}
