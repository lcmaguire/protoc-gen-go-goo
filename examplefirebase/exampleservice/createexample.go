package exampleservice

import (
	context "context"
	errors "errors"
	"fmt"
	"strings"

	connect_go "github.com/bufbuild/connect-go"
	sample "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
)

// CreateExample implements tutorial.ExampleService.CreateExample.
func (e *ExampleService) CreateExample(ctx context.Context, req *connect_go.Request[sample.SearchRequest]) (*connect_go.Response[sample.SearchResponse], error) {
	auth := req.Header().Get("Authorization")
	splitAuth := strings.Split(auth, "Bearer ")
	if len(splitAuth) < 2 {
		return nil, errors.New("header less than 2")
	}
	idToken := splitAuth[1] // Get token.

	if idToken == "" {
		return nil, errors.New("no auth")
	}

	// would need some mapping to DB (maybe)
	testCollectionPath := "testCollection"
	docRef, writeRes, err := e.firestore.Collection(testCollectionPath).Add(ctx, req)
	if err != nil {
		connect_go.NewError(connect_go.CodeInternal, errors.New("asdsadf"))
	}

	fmt.Println(docRef)
	fmt.Println(writeRes)
	// would need some mapping to Return type (maybe)

	/*
		would need a way to get this done nicely that would work, maybe at.
		token, err := e.auth.VerifyIDToken(ctx, idToken)
		if err != nil {
			log.Fatalf("error verifying ID token: %v\n", err)
		}
	*/

	res := connect_go.NewResponse(&sample.SearchResponse{Name: docRef.Path})
	return res, nil
}
