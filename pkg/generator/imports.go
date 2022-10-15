package generator

import "google.golang.org/protobuf/compiler/protogen"

// would be nicer to just have templates and pass in these as strings. Will do that in next piece of work / have these things loaded in via config
var _serviceImports = []protogen.GoImportPath{"log", "net", "google.golang.org/grpc", "google.golang.org/grpc/reflection"}

var _methodImports = []protogen.GoImportPath{"context", "google.golang.org/grpc/codes", "google.golang.org/grpc/status"}

var _testImports = []protogen.GoImportPath{"testing", "context", "testing", "google.golang.org/grpc/codes", "google.golang.org/grpc/status", "github.com/stretchr/testify/assert", "github.com/golang/protobuf/proto"}
