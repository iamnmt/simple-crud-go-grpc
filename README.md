## Prerequisites

- make
- Go
- Protocol buffer compiler
- Go plugins

```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Generate gRPC code

```bash
$ make gen-proto
```

## Run the server

```
$ make run-server
```

## Run the client

```
$ make run-client
```
