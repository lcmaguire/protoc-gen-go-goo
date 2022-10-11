package connectgo

// service appears to be the same

// rpc methods recieve/return connect.Request, connect.NewResponse

// server is much simpler

// import connectgo

import (
	"google.golang.org/protobuf/compiler/protogen"
)

func GenerateFilesForService(gen *protogen.Plugin, service *protogen.Service, file *protogen.File) *protogen.GeneratedFile {
	// need to abstract stuff away, will do later
	generateServiceFile(gen, service)

	for _, v := range service.Methods {

		genRpcMethod(gen, service, v)
		//outfiles = append(outfiles, g)

		// wil generate test file
		//genTestFile(gen, service, v)
		//outfiles = append(outfiles, gT)
	}

	return nil
}

func ConnectGen(gen *protogen.Plugin) {
	//
	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}

		for _, v := range f.Services {
			GenerateFilesForService(gen, v, f)
		}

		GenConnectServer(gen, f)
	}

}
