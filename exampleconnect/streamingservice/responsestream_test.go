package streamingservice

import (
	context "context"
	"net/http"
	"net/http/httptest"
	testing "testing"

	"github.com/bufbuild/connect-go"
	connect_go "github.com/bufbuild/connect-go"
	sample "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
	sampleconnect "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sampleconnect"
	assert "github.com/stretchr/testify/assert"
)

func TestResponseStream(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()

	mux.Handle(sampleconnect.NewStreamingServiceHandler(&StreamingService{}))
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	defer server.Close()

	connectClient := sampleconnect.NewStreamingServiceClient(
		server.Client(),
		server.URL,
	)
	grpcClient := sampleconnect.NewStreamingServiceClient(
		server.Client(),
		server.URL,
		connect_go.WithGRPC(),
	)
	clients := []sampleconnect.StreamingServiceClient{connectClient, grpcClient}

	t.Run("response_stream", func(t *testing.T) {
		total := 0
		for _, client := range clients {
			request := connect.NewRequest(&sample.GreetRequest{})
			stream, err := client.ResponseStream(context.Background(), request)
			assert.Nil(t, err)
			for stream.Receive() {
				total++
			}
			assert.Nil(t, err)
			assert.Error(t, stream.Err())
			assert.Nil(t, stream.Close())
			assert.Equal(t, connect_go.CodeUnimplemented, connect_go.CodeOf(stream.Err()))
			assert.True(t, total > 0)
		}
	})
}
