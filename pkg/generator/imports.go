package generator

import "google.golang.org/protobuf/compiler/protogen"

func standardImports() {

}

// would be nicer to just have templates. Will do that in next piece of work
var _serviceImports = []protogen.GoImportPath{"log", "net", "google.golang.org/grpc", "google.golang.org/grpc/reflection"}
