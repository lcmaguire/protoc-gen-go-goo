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

## Streaming curls.
grpc-connect-streaming: # grpc curl for connect streaming service
	grpcurl \
    -import-path ./exampleconnect -proto example.proto -plaintext \
    -d '{}' \
    localhost:8080 tutorial.StreamingService/ResponseStream


## firebase grpcurls

grpc-firebase-create: 
	grpcurl \
    -import-path ./examplefirebase -proto firebase.proto -plaintext \
    -d '{"name": "testCollection/myname", "display_name": "displayName"}' \
    localhost:8080 tutorial.ExampleService/CreateExample

grpc-firebase-get:
	grpcurl \
    -import-path ./examplefirebase -proto firebase.proto -plaintext \
    -d '{"name": "testCollection/myname"}' \
    localhost:8080 tutorial.ExampleService/GetExample

grpc-firebase-get-auth:
	grpcurl \
    -import-path ./examplefirebase -proto firebase.proto -plaintext \
	-H "Authorization: Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6ImQ3YjE5MTI0MGZjZmYzMDdkYzQ3NTg1OWEyYmUzNzgzZGMxYWY4OWYiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoibGlhbSBlbCBhdXN0cmFsaWFubyIsInBpY3R1cmUiOiJodHRwczovL2xoMy5nb29nbGV1c2VyY29udGVudC5jb20vYS0vQU9oMTRHaUFjRWV4dEg4UXlRalgySjJXSGJITFItZU45TEdSWHRpa0Z3Qmg9czk2LWMiLCJzdG9yZXMiOnsic3RvcmVOYW1lIjoidGVzdCIsImFkbWluIjp0cnVlfSwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL3BiY2MtdGVzdC1lbnYiLCJhdWQiOiJwYmNjLXRlc3QtZW52IiwiYXV0aF90aW1lIjoxNjY4Mzc3MTc0LCJ1c2VyX2lkIjoiSlVOQ3RUa1BJN1oyazUzeE9obTRuSEs3ekJvMiIsInN1YiI6IkpVTkN0VGtQSTdaMms1M3hPaG00bkhLN3pCbzIiLCJpYXQiOjE2NjgzNzcxNzQsImV4cCI6MTY2ODM4MDc3NCwiZW1haWwiOiJtYWduYWxkaW5vcmluaG9AZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZ29vZ2xlLmNvbSI6WyIxMTQ0MDUxNTYyNjI2MTg1NTc3NDgiXSwiZW1haWwiOlsibWFnbmFsZGlub3JpbmhvQGdtYWlsLmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6Imdvb2dsZS5jb20ifX0.CJRqdmQMuhY8OZqeRede3z06jUw1FmB9-VvKGAWBiDAiUUblKI6_u3Ut9oyRfuZiaMyJu9v31Bg4Qxy8rdw3EuDKeJEy4H4ZIv6SWmoJufaDHYVjeXowLelcOgYtJZD7tYAaNWMUb7bhZhLDH1OkxJqs6cjXqajs5Fj1fJVklQCEoDSr4Qw6_CdWH4WHUK-mqOuWPM4ixAApAqDNKIfOSkCRCkDSR8UtnEemCDH-sbF59E60UW7InvIf_MndhNiEfv8ZBCgU3apajOd9DnUGVryq15TBvL5gTyAELpLWfsjfXguS1Souom0joVumzWtjcR4h-p8r3NxLxypbGmZ0wg" \
    -d '{"name": "testCollection/myname"}' \
    localhost:8080 tutorial.ExampleService/GetExample

grpc-firebase-list: 
	grpcurl \
    -import-path ./examplefirebase -proto firebase.proto -plaintext \
	-d '{ }' \
    localhost:8080 tutorial.ExampleService/ListExamples

grpc-firebase-update: 
	grpcurl \
    -import-path ./examplefirebase -proto firebase.proto -plaintext \
    -d '{"name": "testCollection/myname", "display_name": "updated display name"}' \
    localhost:8080 tutorial.ExampleService/UpdateExample

grpc-firebase-delete:
	grpcurl \
    -import-path ./examplefirebase -proto firebase.proto -plaintext \
    -d '{"name": "testCollection/myname"}' \
    localhost:8080 tutorial.ExampleService/DeleteExample

## normal curls for firebase

curl-firebase-create: # normal curl for connect service
	curl \
    --header "Content-Type: application/json" \
	--header "Authorization: Bearer asdfasdf" \
    --data '{"name": "testCollection/myname", "display_name": "displayName"}' \
    http://localhost:8080/tutorial.ExampleService/CreateExample

curl-firebase-get: # normal curl for connect service
	curl \
    --header "Content-Type: application/json" \
	--header "Authorization: Bearer asdfasdf" \
    --data '{"name": "testCollection/myname"}' \
    http://localhost:8080/tutorial.ExampleService/GetExample

curl-firebase-update: # normal curl for connect service
	curl \
    --header "Content-Type: application/json" \
	--header "Authorization: Bearer asdfasdf" \
    --data '{"name": "testCollection/myname", "display_name": "updated curl example"}' \
    http://localhost:8080/tutorial.ExampleService/UpdateExample

curl-firebase-delete: # normal curl for connect service
	curl \
    --header "Content-Type: application/json" \
	--header "Authorization: Bearer asdfasdf" \
    --data '{"name": "testCollection/myname"}' \
    http://localhost:8080/tutorial.ExampleService/DeleteExample

curl-firebase-list: # normal curl for connect service
	curl \
    --header "Content-Type: application/json" \
	--header "Authorization: Bearer asdfasdf" \
    --data '{}' \
    http://localhost:8080/tutorial.ExampleService/ListExamples


start-with-emulator:
	export FIRESTORE_EMULATOR_HOST=localhost:8090 
	go run examplefirebase/cmd/sample/main.go 

start:
	go run examplefirebase/cmd/sample/main.go 