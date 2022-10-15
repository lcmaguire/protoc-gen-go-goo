package connectgo

import "google.golang.org/protobuf/compiler/protogen"

// would be nicer to just have templates. Will do that in next piece of work
var ServiceImports = []protogen.GoImportPath{"log", "net/http", "golang.org/x/net/http2", "golang.org/x/net/http2/h2c"}
