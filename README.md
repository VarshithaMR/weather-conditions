# Introductions
This microservice calls weather api and gives the weather conditions of a particular place.
It is built with Golang-Gorilla mux and Protobuf.
(Exploring proto-buffers with simple microservice)

# Tech
* goversion : 1.22.0
* OS - macOs
* framework - Gorilla mux
* provider : weather api
* IDE : IntelliJ IDEA

# Highlights
1. Protobuf - libprotoc 25.3
2. gRPC Reverse Proxy
# FAQ
1. Install Go - brew install golang
2. Protobuf compiler Installation - brew install protobuf
    * protoc --version
3. Install Protocol Buffers(Protobuf compiler plugin for go) -  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
4. Install grpc (Protobuf compiler plugin for gRPC-go) : go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
5. Setup protoc in PATH :  export PATH="$PATH:$(go env GOPATH)/bin"
6. Refer : https://grpc.io/docs/languages/go/quickstart/
7. How to generate proto files? in weather-conditions repo, run "protoc --go_out=. --go-grpc_out=. proto/weather-condition-protobuf.proto"
8. Install gRPC-gateway to generate reverse proxy : go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.16.0
9. Install gRPC-gateway for openAPI to generate reverse proxy : go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.19.1
10. Clone googleapis : git clone https://github.com/googleapis/googleapis.git, copy "annotations.proto" file to your service in protocol buffers directory
11. Generate : protoc --proto_path=. --proto_path=proto/googleapis --go_out=. --go-grpc_out=. proto/open-mateo-protobuf.proto

