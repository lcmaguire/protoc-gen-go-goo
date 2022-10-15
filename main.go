package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/lcmaguire/protoc-gen-go-goo/pkg/generator"
)

type config struct {
	Server    bool   `yaml:"server"`
	ConnectGo bool   `yaml:"connectGo"`
	GoModPath string `yaml:"goModPath"`
	// tests
	// files to ignore
	// connect-go ?
	// imports
	// server
}

var cfg *config
var GoModPath = ""

func main() {
	var flags flag.FlagSet
	//value := flags.String("param", "", "")
	//out := flags.String("out", "", "")
	cfg = &config{}

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		// todo have this passed in from config
		g := generator.Generator{
			ConnectGo: false,
			Server:    true,
			GoModPath: "",
			Tests:     true,
		}
		// todo have this be used in the Run func
		return g.Run(gen)
	})
}
