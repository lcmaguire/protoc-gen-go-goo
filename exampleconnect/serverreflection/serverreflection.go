package serverreflection

import (
	grpc_reflection_v1alpha "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

// TSUUUU

// ServerReflection ...
type ServerReflection struct {
	grpc_reflection_v1alpha.UnimplementedServerReflectionServer
}
