# run protoc without plugins.
gen: 
	protoc -I=. \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	*.proto

gen-goo:
	go install . && \
	protoc -I=. \
	--go-goo_out=. \
	--go-goo_opt=root=github.com/lcmaguire/protoc-gen-go-goo \
	--go-grpc_out=. \
	--go-grpc_out=paths=source_relative \
	--go_out=. \
	*.proto 
