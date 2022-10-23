# generate only goo generated code, can also include --go-goo_opt= (must have already generated grpc and go code)
gen-goo:
	go install . && \
	protoc -I=example \
	--go-goo_out=example \
	--go-goo_opt=tests=true,server=true,connectGo=false \
	example/*.proto 


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
	go install . && buf generate

grpc-curl: # should return unimplemented
	grpcurl -plaintext localhost:8080 tutorial.ExampleService/GetExample

grpc-curl-reflect: # should return endpoints
	grpcurl -plaintext localhost:8080 list

grpc-curl-connect: # grpc curl for connect service
	grpcurl \
    -import-path ./exampleconnect -proto example.proto -plaintext \
    -d '{}' \
    localhost:8080 tutorial.ExampleService/GetExample
	
curl-connect: # normal curl for connect service
	curl \
    --header "Content-Type: application/json" \
    --data '{}' \
    http://localhost:8080/tutorial.ExampleService/GetExample

curl-connect: # normal curl for connect service
	curl \
    --header "Content-Type: application/json" \
    --data '{}' \
    http://localhost:8080/tutorial.ExampleService/GetExample



curl-connect-create: # normal curl for connect service
	curl \
    --header "Content-Type: application/json" \
	--header "Authorization: Bearer asdfasdf" \
    --data '{"name": "my name", "display_name": "displayName"}' \
    http://localhost:8080/tutorial.ExampleService/CreateExample

curl-connect-get: # normal curl for connect service
	curl \
    --header "Content-Type: application/json" \
	--header "Authorization: Bearer asdfasdf" \
    --data '{"name": "testCollection/uCrDRXkFvBbjp93JURzz"}' \
    http://localhost:8080/tutorial.ExampleService/GetExample


curl-connect-delete: # normal curl for connect service
	curl \
    --header "Content-Type: application/json" \
	--header "Authorization: Bearer asdfasdf" \
    --data '{"name": "testCollection/uCrDRXkFvBbjp93JURzz"}' \
    http://localhost:8080/tutorial.ExampleService/DeleteExample

## Streaming curls.
grpc-connect-streaming: # grpc curl for connect streaming service
	grpcurl \
    -import-path ./exampleconnect -proto example.proto -plaintext \
    -d '{}' \
    localhost:8080 tutorial.StreamingService/ResponseStream
