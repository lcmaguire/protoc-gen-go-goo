package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"gopkg.in/yaml.v3"
)

type config struct {
	Name string `yaml:"name"` // test for now until i decide config.
	// tests
	// files to ignore
	// connect-go ?
	// imports
	// server
}

var cfg *config

func main() {
	// cfg = &config{}
	var flags flag.FlagSet
	value := flags.String("param", "", "")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {

		// todo move to func + set up defaults
		if value != nil && *value != "" {
			bytes, err := ioutil.ReadFile(*value)
			if err != nil {
				panic(err)
			}

			err = yaml.Unmarshal(bytes, &cfg)
			if err != nil {
				panic(err)
			}
		}
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			for _, v := range f.Services {
				generateFilesForService(gen, v, f)
			}

			if len(f.Services) > 0 {
				generateServer(gen, f)
			}
		}
		return nil
	})
}

// generateFile generates a _ascii.pb.go file containing gRPC service definitions.
func generateServiceFile(gen *protogen.Plugin, service *protogen.Service) *protogen.GeneratedFile { // consider returning []
	fileName := strings.ToLower(service.GoName + "/" + service.GoName + ".go") // todo format in snakecase
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{service.GoName}}.go
	g := gen.NewGeneratedFile(fileName, protogen.GoImportPath(service.GoName))
	g.P()
	g.P("package ", strings.ToLower(service.GoName))
	g.P()
	//g.P("// test " + cfg.Name)

	rootGoIndent := gen.FilesByPath[service.Location.SourceFile].GoDescriptorIdent // may run into problems depending on how files are set up.
	pkg := getParamPKG(rootGoIndent.GoImportPath.String())

	structString := formatService(string(service.Desc.Name()), pkg)
	_ = g.QualifiedGoIdent(rootGoIndent) // this auto imports too.

	// todo add in template
	//t := template.Must(template.New("letter").Parse(tplate))

	g.P(structString)
	g.P()
	g.P("// test " + gen.Response().String())
	return g
}

func generateFilesForService(gen *protogen.Plugin, service *protogen.Service, file *protogen.File) (outfiles []*protogen.GeneratedFile) {
	serviceFile := generateServiceFile(gen, service)
	outfiles = append(outfiles, serviceFile)

	// will create a method for all services
	for _, v := range service.Methods {

		g := genRpcMethod(gen, service, v)
		outfiles = append(outfiles, g)

		// wil generate test file
		gT := genTestFile(gen, service, v)
		outfiles = append(outfiles, gT)
	}

	return outfiles
}

func genRpcMethod(gen *protogen.Plugin, service *protogen.Service, method *protogen.Method) *protogen.GeneratedFile {
	filename := strings.ToLower(service.GoName + "/" + method.GoName + ".go")
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{method.GoName}}.go
	g := gen.NewGeneratedFile(filename, protogen.GoImportPath(service.GoName))

	methodCaller := genMethodCaller(service.GoName)                                // maybe methodName or methodReciever
	rootGoIndent := gen.FilesByPath[service.Location.SourceFile].GoDescriptorIdent // may run into problems depending on how files are set up.

	// todo create some func with all required pkgs imported as needed
	g.QualifiedGoIdent(rootGoIndent)                                                                // this auto imports too.
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "context", GoName: ""})                       // it would be nice to figure out how to have this not be aliased
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/codes", GoName: ""})  // it would be nice to figure out how to have this not be aliased
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/status", GoName: ""}) // it would be nice to figure out how to have this not be aliased

	rpcfunc := formatMethod(methodCaller, method.GoName, getParamPKG(method.Input.GoIdent.GoImportPath.String())+"."+method.Input.GoIdent.GoName, getParamPKG(method.Output.GoIdent.GoImportPath.String())+"."+method.Output.GoIdent.GoName)

	g.P()
	g.P("package ", strings.ToLower(service.GoName))
	g.P()
	g.P(rpcfunc)

	return g
}

func genTestFile(gen *protogen.Plugin, service *protogen.Service, method *protogen.Method) *protogen.GeneratedFile {
	filename := strings.ToLower(service.GoName + "/" + method.GoName + "_test.go")
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{method.GoName}}.go
	g := gen.NewGeneratedFile(filename, protogen.GoImportPath(service.GoName))

	g.P()
	g.P("package ", strings.ToLower(service.GoName))
	g.P()

	// rootGoIndent := gen.FilesByPath[service.Location.SourceFile].GoDescriptorIdent // may run into problems depending on how files are set up.

	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "testing", GoName: ""})
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "context", GoName: ""})
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "testing", GoName: ""})
	// g.QualifiedGoIdent(rootGoIndent) come back to this later
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: protogen.GoImportPath(service.GoName), GoName: ""}) // it would be nice to figure out how to have this not be aliased
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/codes", GoName: ""})        // it would be nice to figure out how to have this not be aliased
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/status", GoName: ""})       // it would be nice to figure out how to have this not be aliased
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/status", GoName: ""})       // it would be nice to figure out how to have this not be aliased
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "github.com/stretchr/testify/assert", GoName: ""})  // it would be nice to figure out how to have this not be aliased

	serviceFuncName := service.GoName // for if external pkg is wanted (requires a little more effort) strings.ToLower(service.GoName) + "." + service.GoName
	testFile := formatTestFile(method.GoName, serviceFuncName)
	g.P(testFile)
	return g
}

const methodCallerTemplate = `%s *%s`

func genMethodCaller(in string) string {
	return fmt.Sprintf(methodCallerTemplate, strings.ToLower(in[0:1]), in)
}

func getParamPKG(in string) string {
	arr := strings.Split(in, "/")
	return strings.Trim(arr[len(arr)-1], `"`)
}

const methodTemplate = `
	// %s ...
	func (%s) %s (ctx context.Context, in *%s) (out *%s, err error){
		return nil, status.Error(codes.Unimplemented, "yet to be implemented")
	}
`

// move to go template and use write
func formatMethod(methodCaller string, methodName string, requestType string, responseType string) string {
	return fmt.Sprintf(
		methodTemplate,
		methodName,
		methodCaller,
		methodName,
		requestType,
		responseType,
	)
}

const serviceTemplate = `
// %s ...
type %s struct { 
	%s.Unimplemented%sServer
}
	`

func formatService(serviceName string, pkg string) string {
	return fmt.Sprintf(serviceTemplate, serviceName, serviceName, pkg, serviceName)
}

const testFileTemplate = `
	func Test%s(t *testing.T){
		t.Parallel()
		service := &%s{}
		res, err := service.%s(context.Background(), nil)
		assert.Error(t, err)
		assert.Equal(t, codes.Unimplemented, status.Code(err))
		assert.Nil(t, res)
	}
	`

func formatTestFile(method string, service string) string {
	return fmt.Sprintf(testFileTemplate, method, service, method)
}

// add in reflection api
const serverTemplate = `
package main 

func main() {
    if err := run(); err != nil {
        log.Fatal(err)
    }
}

func run() error {
    listenOn := "127.0.0.1:8080"
    listener, err := net.Listen("tcp", listenOn)
    if err != nil {
        return  err 
    }

    server := grpc.NewServer()
    %s.Register%sServer(server, &%s{}) // this would need to be a list or multiple.
	reflection.Register(server) // this should perhaps be optional
    log.Println("Listening on", listenOn)
    if err := server.Serve(listener); err != nil {
        return err 
    }

    return nil
}

`

// need pkg, services,
func generateServer(gen *protogen.Plugin, file *protogen.File) {
	services := []string{}

	for _, v := range file.Services {
		services = append(services, v.GoName) // service.Desc.Name()
	}

	fileName := strings.ToLower("cmd" + "/" + string(file.GoPackageName) + "/" + "main.go")
	g := gen.NewGeneratedFile(fileName, protogen.GoImportPath("."))
	//g := gen.NewGeneratedFile(fileName, file.GoImportPath)

	pkgName := file.GoDescriptorIdent
	g.QualifiedGoIdent(pkgName)
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: file.GoImportPath})
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "log"})
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "net"})
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc"})
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/grpc/reflection"})

	// need to be for loop , hardcoding for now
	serviceName := services[0] //formatService(string(services[0]), pkg)
	//implementedFileName := strings.ToLower(serviceName + "/" + serviceName + ".go") // todo format in snakecase

	//protogen.GoImportPath(service.GoName)
	// gen.FilesByPath
	// get this to be github.com/lcmaguire/protoc-gen-go-goo/{{serviceName}}
	g.QualifiedGoIdent(protogen.GoImportPath(strings.ToLower(serviceName)).Ident(""))
	// gen.FilesByPath
	//protogen.GoImportPath(service.GoName).GoIdent
	// will be in format /{{goo_out_path}}/{{service.GoName}}/{{service.GoName}}.go
	// g := gen.NewGeneratedFile(fileName, protogen.GoImportPath(service.GoName))

	pkg := getParamPKG(file.GoDescriptorIdent.GoImportPath.String())

	g.P(fmt.Sprintf(
		serverTemplate,
		pkg,
		serviceName,
		strings.ToLower(serviceName)+"."+serviceName, // this needs to be path to where server is written
	// this would need to be a list or multiple. , will be // will be in format /{{goo_out_path}}/{{service.GoName}}/{{service.GoName}}.go
	))
}
