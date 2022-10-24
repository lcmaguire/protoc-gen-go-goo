package firebase

const ListEndpointTemplate = `
// {{.MethodName}} implements {{.FullName}}.
func (s *Service) {{.MethodName}}(ctx context.Context, req *connect_go.Request[{{.RequestType}}]) (*connect_go.Response[{{.ResponseType}}], error) {
	docSnaps, err := s.firestore.Collection("testCollection").Documents(ctx).GetAll() // todo get uid from request.
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("error loading response"))
	}

	res := &{{.ResponseType}}{}
	// would want internal message.
	for _, v := range docSnaps {
		if v == nil || v.Data() == nil {
			return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("error loading response"))
		}
		
		if err := v.DataTo(res); err != nil {
			return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("error unable to load response"))
		}
	}

	return connect_go.NewResponse(res), nil
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

	res := &sample.SearchResponse{}
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
		connect_go.NewError(connect_go.CodeInternal, errors.New("unable to delete."))
	}
	
	res := connect_go.NewResponse(req.Msg) // hard coding for now assuming req and res are same and Write is always successful.
	return res, nil
}

`

const DeleteEndpointTemplate = `
// {{.MethodName}} implements {{.FullName}}.
func (s *Service) {{.MethodName}}ctx context.Context, req *connect_go.Request[{{.RequestType}}]) (*connect_go.Response[{{.ResponseType}}], error) {
	_, err := e.firestore.Doc(req.Msg.Name).Delete(ctx)
	if err != nil {
		connect_go.NewError(connect_go.CodeInternal, errors.New("unable to delete."))
	}

	// Should be &emptypb.Empty{}
	return connect_go.NewResponse(&{{.ResponseType}}{}), nil
}
`

const UpdateEndpointTemplate = `
// {{.MethodName}} implements {{.FullName}}.
func (s *Service) {{.MethodName}}ctx context.Context, req *connect_go.Request[{{.RequestType}}]) (*connect_go.Response[{{.ResponseType}}], error) {
	_, err := s.firestore.Doc(req.Msg.Name).Set(ctx, req.Msg) // .Update may be useful with FieldMask.
	if err != nil {
		connect_go.NewError(connect_go.CodeInternal, errors.New("asdsadf"))
	}

	res := connect_go.NewResponse(req.Msg) // hard coding for now assuming req and res are same and Write is always successful.
	return res, nil
}
`

/*
// {{.MethodName}} implements {{.FullName}}.
func ({{.MethodCaller}}) {{.MethodName}}(ctx context.Context, req *connect_go.Request[{{.RequestType}}]) (*connect_go.Response[{{.ResponseType}}], error) {
	res := connect_go.NewResponse(&{{.ResponseType}}{})
	return res, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("not yet implemented"))
}
*/
