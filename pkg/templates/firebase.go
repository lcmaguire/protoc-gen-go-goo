package templates

const FirebaseServer = `
package main
import (
	"context"
	v4 "firebase.google.com/go/v4"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/api/option"
	"log"
	"net/http"
	// your protoPathHere
	"{{.GenImportPath}}connect"
	// your services
	{{.ServiceImports}}
)
func main() {
	mux := http.NewServeMux()
	// The generated constructors return a path and a plain net/http
	// handler.
	mux.Handle(sampleconnect.NewExampleServiceHandler(createNewService()))
	err := http.ListenAndServe(
		"localhost:8080", // auth host users 8080
		// For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
		// avoid x/net/http2 by using http.ListenAndServeTLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: " + err.Error())
}
// createNewService creates a new Service, exampleservice pkg is hard coded for now
func createNewService() *exampleservice.Service {
	opt := option.WithCredentialsFile("your-firebase-service-account.json") // todo have this be env var
	app, err := v4.NewApp(context.Background(), nil, opt)
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
	return exampleservice.NewService(auth, firestore)
}
`

const FirebaseService = `
package {{.GoPkgName}}
import (
	firestore "cloud.google.com/go/firestore"
	auth "firebase.google.com/go/v4/auth"
	{{.Imports}}
)
// Service implements {{.FullName}}.
type Service struct {
	{{.Pkg}}.Unimplemented{{.ServiceName}}Handler
	firestore *firestore.Client
	auth      *auth.Client
}
func NewService(auth *auth.Client, firestore *firestore.Client) *Service {
	return &Service{
		auth:      auth,
		firestore: firestore,
	}
}
`

const FirebaseCreateMethod = `
package {{.GoPkgName}}
import (
	"context"
	connect_go "github.com/bufbuild/connect-go"
	{{.Imports}}
)
// {{.MethodName}} implements {{.FullName}}.
func (s *Service) {{.MethodName}}(ctx context.Context, req *connect_go.Request[{{.RequestType}}]) (*connect_go.Response[{{.ResponseType}}], error) {
	_, err := s.firestore.Doc(req.Msg.Name).Create(ctx, req.Msg)
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}
	res := connect_go.NewResponse(req.Msg) // hard coding for now assuming req and res type are same and Write is always successful.
	return res, nil
}
`

const FirebaseUpdateMethod = `
package {{.GoPkgName}}
import (
	"context"
	connect_go "github.com/bufbuild/connect-go"
	{{.Imports}}
)
// {{.MethodName}} implements {{.FullName}}.
func (s *Service) {{.MethodName}}(ctx context.Context, req *connect_go.Request[{{.RequestType}}]) (*connect_go.Response[{{.ResponseType}}], error) {
	_, err := s.firestore.Doc(req.Msg.Name).Set(ctx, req.Msg) // .Update may be useful with FieldMask.
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}
	res := connect_go.NewResponse(req.Msg) // hard coding for now assuming req and res type are same and Write is always successful.
	return res, nil
}
`

const FirebaseDeleteMethod = `
package {{.GoPkgName}}
import (
	"context"
	connect_go "github.com/bufbuild/connect-go"
	{{.Imports}}
)
// {{.MethodName}} implements {{.FullName}}.
func (s *Service) {{.MethodName}}(ctx context.Context, req *connect_go.Request[{{.RequestType}}]) (*connect_go.Response[{{.ResponseType}}], error) {
	_, err := s.firestore.Doc(req.Msg.Name).Delete(ctx)
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}
	return connect_go.NewResponse(&{{.ResponseType}}{}), nil
}
`

const FirebaseGetMethod = `
package {{.GoPkgName}}
import (
	"context"
	connect_go "github.com/bufbuild/connect-go"
	
	{{.Imports}}
)
// {{.MethodName}} implements {{.FullName}}.
func (s *Service) {{.MethodName}}(ctx context.Context, req *connect_go.Request[{{.RequestType}}]) (*connect_go.Response[{{.ResponseType}}], error) {
	docSnap, err := s.firestore.Doc(req.Msg.Name).Get(ctx)
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}
	if docSnap == nil || docSnap.Data() == nil {
		return nil, connect_go.NewError(connect_go.CodeNotFound, err)
	}
	res := &{{.ResponseType}}{}
	if err := docSnap.DataTo(res); err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}
	return connect_go.NewResponse(res), nil
}
`

const FirebaseListMethod = `
package {{.GoPkgName}}
import (
	"context"
	connect_go "github.com/bufbuild/connect-go"
	
	{{.Imports}}
)
// {{.MethodName}} implements {{.FullName}}.
func (s *Service) {{.MethodName}}(ctx context.Context, req *connect_go.Request[{{.RequestType}}]) (*connect_go.Response[{{.ResponseType}}], error) {
	docSnaps, err := s.firestore.Collection("testCollection").Documents(ctx).GetAll() // hardcoding collection for now. Should probably be MessageName plural.
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}
	arr := make([]*{{.Pkg}}.{{.MessageName}}, 0)
	for _, v := range docSnaps {
		if v == nil || v.Data() == nil {
			return nil, connect_go.NewError(connect_go.CodeInternal, err)
		}
		var data *{{.Pkg}}.{{.MessageName}}
		if err := v.DataTo(&data); err != nil {
			return nil, connect_go.NewError(connect_go.CodeInternal, err)
		}
		arr = append(arr, data)
	}
	return connect_go.NewResponse(
		&{{.ResponseType}}{
			{{.MessageName}}s: arr,
		},
	), nil
}
`
