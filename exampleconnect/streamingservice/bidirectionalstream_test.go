package streamingservice

import (
	context "context"
	errors "errors"
	fmt "fmt"
	connect_go "github.com/bufbuild/connect-go"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
	io "io"
	http "net/http"
	httptest "net/http/httptest"
	sync "sync"
	testing "testing"

	"github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
	"github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sampleconnect"
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

	t.Run("bidirectionalTest", func(t *testing.T) {
		for _, client := range clients {
			sendValues := []string{"Hello!", "How are you doing?", "I have an issue with my bike", "bye"}
			var receivedValues []string
			stream := client.BiDirectionalStream(context.Background())
			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				defer wg.Done()
				for _, sentence := range sendValues {
					err := stream.Send(&sample.GreetRequest{})
					require.Nil(t, err)
					fmt.Println(sentence)
				}
				require.Nil(t, stream.CloseRequest())
			}()
			go func() {
				defer wg.Done()
				for {
					_, err := stream.Receive()
					if errors.Is(err, io.EOF) {
						break
					}
					require.Nil(t, err)
					receivedValues = append(receivedValues, "")
				}
				require.Nil(t, stream.CloseResponse())
			}()
			wg.Wait()
			assert.Equal(t, len(receivedValues), len(sendValues))
		}
	})
}
