package generator

import "google.golang.org/protobuf/compiler/protogen"

// move to be protogen.GoIdent
var _serviceImports = []protogen.GoImportPath{"log", "net", "google.golang.org/grpc", "google.golang.org/grpc/reflection"}

var _methodImports = []protogen.GoImportPath{"context", "google.golang.org/grpc/codes", "google.golang.org/grpc/status"}

var _testImports = []protogen.GoImportPath{"testing", "context", "testing", "google.golang.org/grpc/codes", "google.golang.org/grpc/status", "github.com/stretchr/testify/assert", "github.com/golang/protobuf/proto"}
