package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/lcmaguire/protoc-gen-go-goo/pkg/generator"
)

func main() {
	var flags flag.FlagSet
	connectGo := flags.Bool("connectGo", false, "to generate code for a connect go service, by default it will assume grpc-go")
	tests := flags.Bool("tests", true, "to generate tests for your service")
	server := flags.Bool("server", false, "will generate a basic server that implements your services.")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		g := &generator.Generator{
			ConnectGo: *connectGo,
			Server:    *tests,
			GoModPath: "", // todo implement this so server gen doesn't need to be hardcoded.
			Tests:     *server,
		}
		// todo have this be used in place of the Run func (if possible)
		return g.Run(gen)
	})
}
