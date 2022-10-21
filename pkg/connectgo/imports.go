package connectgo

import "google.golang.org/protobuf/compiler/protogen"

// move to be protogen.GoIdent OR strings
var ServiceImports = []protogen.GoImportPath{"log", "net/http", "golang.org/x/net/http2", "golang.org/x/net/http2/h2c"}

var MethodImports = []protogen.GoImportPath{"context", "github.com/bufbuild/connect-go", "errors"}

var TestImports = []protogen.GoImportPath{"context", "github.com/bufbuild/connect-go", "testing", "testing", "github.com/stretchr/testify/assert", "github.com/golang/protobuf/proto"}
