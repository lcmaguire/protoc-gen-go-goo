package exampleservice

import (
	context "context"
	"fmt"
	"log"
	testing "testing"

	firebase "firebase.google.com/go/v4"
	connect_go "github.com/bufbuild/connect-go"
	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
	assert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/api/option"
)

func TestCreateExample(t *testing.T) {
	t.Parallel()
	service := createNewService()
	req := &connect_go.Request[sample.SearchRequest]{
		Msg: &sample.SearchRequest{
			Name: "res name",
		},
	}
	ctx := context.Background()
	uid := "JUNCtTkPI7Z2k53xOhm4nHK7zBo2"
	tok, err := service.auth.CustomToken(ctx, uid)
	//fmt.Println("antes")
	require.NoError(t, err)
	// fmt.Println("despues")
	req.Header().Set("Authorization", "Bearer "+tok)

	//u, err := service.auth.GetUser(ctx, uid)
	//require.NoError(t, err)

	res, err := service.CreateExample(context.Background(), req)
	assert.NoError(t, err)
	fmt.Println(res)

}

func createNewService() *ExampleService {
	opt := option.WithCredentialsFile("./../../test-firebase-service-account.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	firestore, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return NewExampleService(auth, firestore)
}
