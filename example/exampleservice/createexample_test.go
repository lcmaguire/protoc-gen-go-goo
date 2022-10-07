package exampleservice

import (
	context "context"
	assert "github.com/stretchr/testify/assert"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	testing "testing"
)

func TestCreateExample(t *testing.T) {
	t.Parallel()
	service := &ExampleService{}
	res, err := service.CreateExample(context.Background(), nil)
	assert.Error(t, err)
	assert.Equal(t, codes.Unimplemented, status.Code(err))
	assert.Nil(t, res)
}
