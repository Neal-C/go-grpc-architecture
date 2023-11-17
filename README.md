# A Go Playground

## Calm down, this is not meant for production. it is not "prod ready"

### Technical & design choices

Every design and tech decision taken, were based, in order of importance by: what I want to experiment/discover, what I like, what's been recommended to me by the Go community

### Tech Stack

- gRPC
- protobuf
- google grpc

### Requirements

- protoc version >=3

- Go protocol compiler plugins

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-rpc@v1.2
```