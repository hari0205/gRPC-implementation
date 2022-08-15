
# gRPC template for Go

A template for running gRPC implementations. 


## Prerequisites

- [Go](https://go.dev/dl/)




## Installation

- Clone this repository

- Run the following to install protocol compiler plugins
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

```

- Update GO ENV path:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"

``` 

- Edit the `.proto` file as per your liking

- Run the following command
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/*.proto

```

- Add implementations to `client.go` and `server.go`

- Finally , Run
```
go run server/server.go
go run client/client.go

```
