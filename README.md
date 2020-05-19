## Installation
**gen-go**

`$ go get github.com/golang/protobuf/protoc-gen-go`

`$ go install github.com/golang/protobuf/protoc-gen-go`

**grpc-gateway**

`$ go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway`

`$ go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway`

`$ go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger`

`$ go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger`





## Gen Command
**Go Code**

`$ protoc --proto_path=./pb --go_out=plugins=grpc:. pb/hello.proto`

**grpc-gateway**

`$ protoc -I=%GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I=. -I=%GOPATH%/src --grpc-gateway_out=logtostderr=true,paths=source_relative:. pb/hello.proto`

**swagger**

`$ protoc -I=%GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I=. -I=%GOPATH%/src --swagger_out=logtostderr=true:. pb/hello.proto`


## Ref
[grpc.io](https://grpc.io/)

[grpc-gateway](https://grpc-ecosystem.github.io/grpc-gateway/docs/usage.html)