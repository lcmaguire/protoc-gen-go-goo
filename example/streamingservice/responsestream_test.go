package streamingservice

import (
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
	assert "github.com/stretchr/testify/assert"
	testing "testing"
)

func TestResponseStream(t *testing.T) {
	t.Parallel()
	// tests for this type of RPC yet to be implemented.
	assert.NotNil(t, &example.GreetRequest{})
	assert.NotNil(t, &example.GreetResponse{})
}
