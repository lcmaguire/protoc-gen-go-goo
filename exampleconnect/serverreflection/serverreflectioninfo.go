package serverreflection

import (
	context "context"
	connect_go "github.com/bufbuild/connect-go"
	codes "google.golang.org/grpc/codes"
	grpc_reflection_v1alpha "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	status "google.golang.org/grpc/status"
)

func (s *ServerReflection) ServerReflectionInfo(ctx context.Context, req *connect_go.Request[grpc_reflection_v1alpha.ServerReflectionRequest]) (*connect_go.Response[grpc_reflection_v1alpha.ServerReflectionResponse], error) {
	res := connect_go.NewResponse(&grpc_reflection_v1alpha.ServerReflectionResponse{})
	return res, status.Error(codes.Unimplemented, "yet to be implemented")
}
