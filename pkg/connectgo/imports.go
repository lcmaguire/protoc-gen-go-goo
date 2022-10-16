package connectgo

import "google.golang.org/protobuf/compiler/protogen"

// move to be protogen.GoIdent OR strings
var ServiceImports = []protogen.GoImportPath{"log", "net/http", "golang.org/x/net/http2", "golang.org/x/net/http2/h2c"}

var MethodImports = []protogen.GoImportPath{"context", "google.golang.org/grpc/codes", "google.golang.org/grpc/status", "github.com/bufbuild/connect-go"}

var TestImports = append(MethodImports, "testing", "testing", "github.com/stretchr/testify/assert", "github.com/golang/protobuf/proto")
