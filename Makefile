# generates grpc-go, go proto code and goo generated code, can also include -go-goo_opt=
grpc-go-goo:
	go install . && \
	protoc -I=example \
	--go-goo_out=example \
	--go-goo_opt=tests=true,server=true,connectGo=false,generatedPath=github.com/lcmaguire/protoc-gen-go-goo/example \
	--go-grpc_out=example \
	--go-grpc_opt=paths=source_relative \
	--go_out=example  \
	--go_opt=paths=source_relative \
	example/*.proto 

# generates connect-go and go-proto code 
connect-go-goo:
	go install . && \
	protoc -I=exampleconnect \
	--go-goo_out=exampleconnect \
	--go-goo_opt=tests=true,server=true,connectGo=true,generatedPath=github.com/lcmaguire/protoc-gen-go-goo/exampleconnect \
	--go_out=exampleconnect/sample  \
	--go_opt=paths=source_relative \
	--connect-go_out=exampleconnect \
	--connect-go_opt=paths=source_relative  \
	exampleconnect/example.proto 

# same as above but easier to manage via buf files.
make buf:
	go install 
	buf generate

# same as above but easier to manage via buf files.
make buf-connect:
	go install 
	buf generate --template buf.gen.connect.yaml
	
grpc-curl: # should return unimplemented
	grpcurl -plaintext localhost:8080 tutorial.ExampleService/GetExample

grpc-curl-reflect: # should return endpoints
	grpcurl -plaintext localhost:8080 list

grpc-curl-connect: # grpc curl for connect service
	grpcurl \
    -import-path ./examplefirebase -proto example.proto -plaintext \
    -d '{}' \
    localhost:8080 tutorial.ExampleService/ListExamples
	
curl-connect: # normal curl for connect service
	curl \
    --header "Content-Type: application/json" \
    --data '{}' \
    http://localhost:8080/tutorial.ExampleService/ListExamples

## Streaming curls.
grpc-connect-streaming: # grpc curl for connect streaming service
	grpcurl \
    -import-path ./exampleconnect -proto streaming.proto -plaintext \
    -d '{}' \
    localhost:8080 tutorial.StreamingService/ResponseStream

run-firebase:
	go run ./examplefirebase/cmd/sample/main.go