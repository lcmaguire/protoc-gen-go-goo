package streamingservice

import (
	context "context"
	"fmt"
	"net/http"
	"net/http/httptest"
	testing "testing"

	"github.com/bufbuild/connect-go"
	connect_go "github.com/bufbuild/connect-go"

	sample "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
	sampleconnect "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sampleconnect"
	assert "github.com/stretchr/testify/assert"
)

func TestBiDirectionalStream(t *testing.T) {
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
		for _, client := range clients {
			result, err := client.ResponseStream(context.Background(), connect.NewRequest(&sample.GreetRequest{}))
			assert.NoError(t, err)
			assert.NotNil(t, result)
			fmt.Println(err)
		}
	})
}
