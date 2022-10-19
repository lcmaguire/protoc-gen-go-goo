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
	--go-goo_opt=tests=true,server=true,connectGo=false \
	--go-grpc_out=example \
	--go-grpc_opt=paths=source_relative \
	--go_out=example  \
	--go_opt=paths=source_relative \
	example/*.proto 

# generates grpc-go and go-proto code 
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

## try -> move codegen to gen/ link in https://connect.build/docs/go/getting-started/
## maybe connectexample/gen

#gooey/yeet
connect-go-goo-diff:
	go install . && \
	protoc -I=example \
	--go-goo_out=../gooey \
	--go-goo_opt=tests=true,server=true,connectGo=true,generatedPath=gooey \
	--go_out=exampleconnect  \
	--go_opt=paths=source_relative \
	--connect-go_out=exampleconnect \
	--connect-go_opt=paths=source_relative  \
	example/*.proto 

# same as above but easier to manage via buf files.
make buf:
	go install . && buf generate

grpc-curl: # should return unimplemented
	grpcurl -plaintext localhost:8080 tutorial.ExampleService/GetExample

grpc-curl-reflect: # should return endpoints
	grpcurl -plaintext localhost:8080 list

grpc-curl-connect: # grpc curl for connect service
	grpcurl \
    -import-path ./example -proto example.proto -plaintext \
    -d '{}' \
    localhost:8080 tutorial.ExampleService/GetExample

curl-connect: # normal curl for connect service
	curl \
    --header "Content-Type: application/json" \
    --data '{}' \
    http://localhost:8080/tutorial.ExampleService/GetExample

