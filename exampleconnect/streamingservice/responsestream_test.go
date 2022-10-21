package streamingservice

import (
	sample "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
	assert "github.com/stretchr/testify/assert"
	testing "testing"
)

func TestResponseStream(t *testing.T) {
	t.Parallel()
	// tests for this type of RPC yet to be implemented.
	assert.NotNil(t, &sample.GreetRequest{})
	assert.NotNil(t, &sample.GreetResponse{})
}
