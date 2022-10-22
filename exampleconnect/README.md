the following command when ran from the root of this project can be used to generate the code and files in this directory

```
go install . && \
	protoc -I=exampleconnect \
	--go-goo_out=exampleconnect \
	--go-goo_opt=tests=true,server=true,connectGo=true,generatedPath=github.com/lcmaguire/protoc-gen-go-goo/exampleconnect \
	--go_out=exampleconnect/sample  \
	--go_opt=paths=source_relative \
	--connect-go_out=exampleconnect \
	--connect-go_opt=paths=source_relative  \
	exampleconnect/example.proto 

```