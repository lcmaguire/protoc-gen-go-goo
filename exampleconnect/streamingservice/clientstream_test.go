package streamingservice

import (
	"context"
	connect_go "github.com/bufbuild/connect-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	sample "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
	"github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sampleconnect"
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
