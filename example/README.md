the following command when ran from the root of this project can be used to generate the code and files in this directory

```
# generates grpc-go, go proto code and goo generated code, can also include -go-goo_opt=

go install . && \
	protoc -I=example \
	--go-goo_out=example \
	--go-goo_opt=tests=true,server=true,connectGo=false,generatedPath=github.com/lcmaguire/protoc-gen-go-goo/example \
	--go-grpc_out=example \
	--go-grpc_opt=paths=source_relative \
	--go_out=example  \
	--go_opt=paths=source_relative \
	example/*.proto 

```