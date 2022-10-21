package streamingservice

import (
	testing "testing"

	sample "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
	assert "github.com/stretchr/testify/assert"
)

func TestClientStream(t *testing.T) {
	t.Parallel()
	// tests for this type of RPC yet to be implemented.
	assert.NotNil(t, &sample.GreetRequest{})
	assert.NotNil(t, &sample.GreetResponse{})

	/*
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
	*/

	/*
		t.Run("introduce", func(t *testing.T) {
			total := 0
			for _, client := range clients {
				request := connect.NewRequest(&sample.SearchRequest{})
				stream, err := client.ClientStream(context.Background(), request)
				assert.Nil(t, err)
				for stream.Receive() {
					total++
				}
				assert.Nil(t, stream.Err())
				assert.Nil(t, stream.Close())
				assert.True(t, total > 0)
			}
		})*/
}
