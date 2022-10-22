package connectgo

import "google.golang.org/protobuf/compiler/protogen"

var ServiceImports = []protogen.GoImportPath{"log", "net/http", "golang.org/x/net/http2", "golang.org/x/net/http2/h2c"}

var MethodImports = []protogen.GoImportPath{"context", "github.com/bufbuild/connect-go", "errors"}

var TestImports = []protogen.GoImportPath{"context", "github.com/bufbuild/connect-go", "testing", "testing", "github.com/stretchr/testify/assert", "github.com/golang/protobuf/proto"}

var TestBiDirectionalMethod = []protogen.GoImportPath{"testing", "context", "github.com/bufbuild/connect-go", "errors", "fmt", "io", "net/http", "net/http/httptest", "sync", "github.com/stretchr/testify/assert", "github.com/stretchr/testify/require"}

var TestClientStreamMethod = []protogen.GoImportPath{"testing", "context", "github.com/bufbuild/connect-go", "net/http", "net/http/httptest", "sync", "github.com/stretchr/testify/assert", "github.com/stretchr/testify/require"}

var TestServerStreamMethod = []protogen.GoImportPath{"testing", "context", "github.com/bufbuild/connect-go", "net/http", "net/http/httptest", "github.com/stretchr/testify/assert"}
