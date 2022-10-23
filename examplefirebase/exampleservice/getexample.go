package exampleservice

import (
	context "context"
	"encoding/json"
	errors "errors"
	"fmt"

	connect_go "github.com/bufbuild/connect-go"
	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
	"google.golang.org/protobuf/encoding/protojson"
)

// GetExample implements tutorial.ExampleService.GetExample.
func (e *ExampleService) GetExample(ctx context.Context, req *connect_go.Request[sample.SearchRequest]) (*connect_go.Response[sample.SearchResponse], error) {
	docSnap, err := e.firestore.Doc(req.Msg.Name).Get(ctx)
	if err != nil {
		connect_go.NewError(connect_go.CodeInternal, errors.New("asdsadf"))
	}

	if docSnap == nil || docSnap.Data() == nil {
		return nil, connect_go.NewError(connect_go.CodeNotFound, errors.New("err not found"))
	}

	data := docSnap.Data()
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("err marshal"))
	}

	// proto.Marshal()
	res := &sample.SearchResponse{}
	if err := protojson.Unmarshal(bytes, res); err != nil {
		fmt.Println(data)
		return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("unmarshal"))
	}

	return connect_go.NewResponse(res), nil
}
