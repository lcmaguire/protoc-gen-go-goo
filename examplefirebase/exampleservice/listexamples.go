package exampleservice

import (
	context "context"
	errors "errors"

	connect_go "github.com/bufbuild/connect-go"
	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

// ListExamples implements tutorial.ExampleService.ListExamples.
func (s *Service) ListExamples(ctx context.Context, req *connect_go.Request[sample.ListExampleRequest]) (*connect_go.Response[sample.ListExampleResponse], error) {
	// todo have some form of control over collections.
	docSnaps, err := s.firestore.Collection("testCollection").Documents(ctx).GetAll() // todo get uid from request.
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("error loading response"))
	}

	res := &sample.ListExampleResponse{
		Examples: []*sample.Example{},
	}
	// {{MessageType}}s: []*{{pkg}}.{{MessageTypeStripped}}{}
	// would want internal message.
	for _, v := range docSnaps {
		if v == nil || v.Data() == nil {
			return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("error loading response"))
		}

		// var data *{{pkg}}.{{MessageType}}
		var data *sample.Example
		if err := v.DataTo(data); err != nil {
			return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("error unable to load response"))
		}
		// res.{{messageType}}s = append(res.{{messageType}}s, data)
		res.Examples = append(res.Examples, data)
	}

	return connect_go.NewResponse(res), nil
}
