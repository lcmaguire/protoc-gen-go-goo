package connectgo

import "google.golang.org/protobuf/compiler/protogen"

// move to be protogen.GoIdent OR strings
var ServiceImports = []protogen.GoImportPath{"log", "net/http", "golang.org/x/net/http2", "golang.org/x/net/http2/h2c"}

var MethodImports = []protogen.GoImportPath{"context", "github.com/bufbuild/connect-go", "errors"}

var TestImports = []protogen.GoImportPath{"context", "github.com/bufbuild/connect-go", "testing", "testing", "github.com/stretchr/testify/assert", "github.com/golang/protobuf/proto"}

var TestBiDirectionalMethod = []protogen.GoImportPath{"testing", "context", "github.com/bufbuild/connect-go", "errors", "fmt", "io", "net/http", "net/http/httptest", "sync", "github.com/stretchr/testify/assert", "github.com/stretchr/testify/require"}

/*

context "context"
	errors "errors"
	"fmt"
	io "io"
	"net/http"
	"net/http/httptest"
	"sync"
	testing "testing"

	connect_go "github.com/bufbuild/connect-go"

	sample "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
	sampleconnect "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sampleconnect"
	assert "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

*/
