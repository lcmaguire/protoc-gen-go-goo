# generate only goo generated code, can also include --go-goo_opt=
gen-goo:
	go install . && \
	protoc -I=example \
	--go-goo_out=. \
	*.proto 

# generate only goo generated code, can also include -go-goo_opt=
gen-goo-config:
	go install . && \
	protoc -I=example \
	--go-goo_out=example \
	--go-goo_opt=param=config.yaml \
	example/*.proto

# generate grpc-go, and go-proto code + goo generated code, can also include -go-goo_opt=
gen-goo-proto-too:
	go install . && \
	protoc -I=example \
	--go-goo_out=example \
	--go-grpc_out=example \
	--go-grpc_opt=paths=source_relative \
	--go_out=example  \
	--go_opt=paths=source_relative \
	example/*.proto 

grpc-curl: # should return unimplemented
	grpcurl -plaintext localhost:8080 tutorial.ExampleService/GetExample

grpc-curl-reflect: # should return endpoints
	grpcurl -plaintext localhost:8080 list

grpc-curl-connect:
	grpcurl \
    -import-path ./example -proto example.proto -plaintext \
    -d '{}' \
    localhost:8080 tutorial.ExampleService/GetExample

curl-connect:
	curl \
    --header "Content-Type: application/json" \
    --data '{}' \
    http://localhost:8080/tutorial.ExampleService/GetExample

grpc-protoc-gen-connect-go:
	go install . && \
	protoc -I=example \
	--go-goo_out=exampleconnect \
	--go-goo_opt=param=config.yaml \
	--go_out=exampleconnect  \
	--go_opt=paths=source_relative \
	--connect-go_out=exampleconnect \
	--connect-go_opt=paths=source_relative  \
	example/*.proto 

make buf:
	go install . && buf generate
