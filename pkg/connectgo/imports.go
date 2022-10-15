package connectgo

import "google.golang.org/protobuf/compiler/protogen"

// would be nicer to just have templates. Will do that in next piece of work
var ServiceImports = []protogen.GoImportPath{"log", "net/http", "golang.org/x/net/http2", "golang.org/x/net/http2/h2c"}

var TestImports = []protogen.GoImportPath{"testing", "context", "testing", "google.golang.org/grpc/codes", "google.golang.org/grpc/status", "github.com/stretchr/testify/assert"}

var MethodImports = []protogen.GoImportPath{"context", "google.golang.org/grpc/codes", "google.golang.org/grpc/status", "github.com/bufbuild/connect-go"}

/*

	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "context", GoName: ""})
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/codes", GoName: ""})
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/status", GoName: ""})
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "github.com/bufbuild/connect-go", GoName: ""})

*/
