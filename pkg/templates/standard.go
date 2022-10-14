package templates

const serviceTemplate = `
// %s ...
type %s struct { 
	%s.Unimplemented%sServer
}
	`

// add in reflection api
const serverTemplate = `
func main() {
    if err := run(); err != nil {
        log.Fatal(err)
    }
}

func run() error {
    listenOn := "127.0.0.1:8080" // this should be passed in via config
    listener, err := net.Listen("tcp", listenOn) // this too
    if err != nil {
        return  err 
    }

    server := grpc.NewServer()
	// services in your protoFile
    %s
	log.Println("Listening on", listenOn)
    if err := server.Serve(listener); err != nil {
        return err 
    }

    return nil
}

`

const registerServiceTemplate = `
%s.Register%sServer(server, &%s{})
reflection.Register(server) // this should perhaps be optional

`
