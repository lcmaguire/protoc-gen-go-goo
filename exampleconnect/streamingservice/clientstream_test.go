package streamingservice

import (
	context "context"
	connect_go "github.com/bufbuild/connect-go"
	sample "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
	sampleconnect "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sampleconnect"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	http "net/http"
	httptest "net/http/httptest"
	sync "sync"
	testing "testing"
)

func TestClientStream(t *testing.T) {
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

	for _, client := range clients {
		sendValues := []string{"Hello!", "How are you doing?", "I have an issue with my bike", "bye"}
		stream := client.ClientStream(context.Background())
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, sentence := range sendValues {
				err := stream.Send(&sample.GreetRequest{})

				require.Nil(t, err, sentence)
			}
		}()
		wg.Wait()
		res, err := stream.CloseAndReceive()
		require.Error(t, err)
		assert.Equal(t, connect_go.CodeUnimplemented, connect_go.CodeOf(err))
		assert.Nil(t, res)
	}
}
