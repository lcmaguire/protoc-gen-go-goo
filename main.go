package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/lcmaguire/protoc-gen-go-goo/pkg/generator"
)

func main() {
	var flags flag.FlagSet
	connectGo := flags.Bool("connectGo", false, "to generate code for a connect go service, by default it will assume grpc-go")
	tests := flags.Bool("tests", true, "to generate tests for your service")
	server := flags.Bool("server", false, "will generate a basic server that implements your services.")
	generatedGoModPath := flags.String("generatedPath", "", "GoModPath + generated code path so that server will import correctly.")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		g := &generator.Generator{
			ConnectGo: *connectGo,
			Server:    *server,
			GoModPath: *generatedGoModPath, // todo implement this so server gen doesn't need to be hardcoded.
			Tests:     *tests,
		}
		// this enables optional fields to be supported.
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		// todo have this be used in place of the Run func (if possible)
		return g.Run(gen)
	})
}
