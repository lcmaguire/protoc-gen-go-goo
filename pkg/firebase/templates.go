package firebase

const ListEndpointTemplate = `
// {{.MethodName}} implements {{.FullName}}.
func (s *Service) {{.MethodName}}(ctx context.Context, req *connect_go.Request[{{.RequestType}}]) (*connect_go.Response[{{.ResponseType}}], error) {
	docSnaps, err := s.firestore.Collection("testCollection").Documents(ctx).GetAll() // todo get uid from request.
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("error loading response"))
	}
	arr := []*{{.ProtoPkg}}.{{.MessageName}}{}
	for _, v := range docSnaps {
		if v == nil || v.Data() == nil {
			return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("error loading response"))
		}
		
		var data *{{.ProtoPkg}}.{{.MessageName}}
		if err := v.DataTo(&data); err != nil {
			return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("error unable to load response"))
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

const GetEndpointTemplate = `
// {{.MethodName}} implements {{.FullName}}.
func (s *Service) {{.MethodName}}(ctx context.Context, req *connect_go.Request[{{.RequestType}}]) (*connect_go.Response[{{.ResponseType}}], error) {
	docSnap, err := s.firestore.Doc(req.Msg.Name).Get(ctx)
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("error loading response"))
	}

	if docSnap == nil || docSnap.Data() == nil {
		return nil, connect_go.NewError(connect_go.CodeNotFound, errors.New("err not found"))
	}

	res := &{{.ResponseType}}{}
	if err := docSnap.DataTo(res); err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("error loading response"))
	}
	return connect_go.NewResponse(res), nil
}

`

const CreateEndpointUpdate = `
// {{.MethodName}} implements {{.FullName}}.
func (s *Service) {{.MethodName}}(ctx context.Context, req *connect_go.Request[{{.RequestType}}]) (*connect_go.Response[{{.ResponseType}}], error) {
	_, err := s.firestore.Doc(req.Msg.Name).Create(ctx, req.Msg)
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("error loading response"))
	}
	
	res := connect_go.NewResponse(req.Msg) // hard coding for now assuming req and res are same and Write is always successful.
	return res, nil
}

`

const DeleteEndpointTemplate = `
// {{.MethodName}} implements {{.FullName}}.
func (s *Service) {{.MethodName}}(ctx context.Context, req *connect_go.Request[{{.RequestType}}]) (*connect_go.Response[{{.ResponseType}}], error) {
	_, err := s.firestore.Doc(req.Msg.Name).Delete(ctx)
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("error loading response"))
	}

	// Should be &emptypb.Empty{}
	return connect_go.NewResponse(&{{.ResponseType}}{}), nil
}
`

const UpdateEndpointTemplate = `
// {{.MethodName}} implements {{.FullName}}.
func (s *Service) {{.MethodName}}(ctx context.Context, req *connect_go.Request[{{.RequestType}}]) (*connect_go.Response[{{.ResponseType}}], error) {
	_, err := s.firestore.Doc(req.Msg.Name).Set(ctx, req.Msg) // .Update may be useful with FieldMask.
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("error loading response"))
	}

	res := connect_go.NewResponse(req.Msg) // hard coding for now assuming req and res are same and Write is always successful.
	return res, nil
}
`

const ServiceTemplate = `
	// Service implements {{.FullName}}.
	type Service struct {
		sampleconnect.Unimplemented{{.ServiceName}}Handler
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

const ServerTemplate = `

func main() {
	mux := http.NewServeMux()
	// The generated constructors return a path and a plain net/http
	// handler.

	mux.Handle(sampleconnect.NewExampleServiceHandler(createNewService()))

	// export FIREBASE_AUTH_EMULATOR_HOST="localhost:9099"
	// export FIRESTORE_EMULATOR_HOST="localhost:8080"

	err := http.ListenAndServe(
		"localhost:8080", // auth host users 8080
		// For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
		// avoid x/net/http2 by using http.ListenAndServeTLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: " + err.Error())
}

func createNewService() *exampleservice.Service {
	opt := option.WithCredentialsFile("./test-firebase-service-account.json")
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
