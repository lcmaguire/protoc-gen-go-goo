# generate only goo generated code, can also include --go-goo_opt=
gen-goo:
	go install . && \
	protoc -I=. \
	--go-goo_out=. \
	*.proto 

# generate only goo generated code, can also include -go-goo_opt=
gen-goo-config:
	go install . && \
	protoc -I=. \
	--go-goo_out=. \
	--go-goo_opt=param=config.yaml \
	*.proto 

# generate grpc-go, and go-proto code + goo generated code, can also include -go-goo_opt=
gen-goo-proto-too:
	go install . && \
	protoc -I=. \
	--go-goo_out=. \
	--go-grpc_out=example/out \
	--go-grpc_opt=paths=source_relative \
	--go_out=example/out  \
	--go_opt=paths=source_relative \
	*.proto 
