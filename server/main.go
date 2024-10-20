package main

import (
	"log"
	"net"

	"ms/server/handler"
	"ms/server/service"
	"google.golang.org/grpc"
)

const (
	addr = "localhost:6969"
)

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed with %v", err)
	}

	grpcServer := grpc.NewServer()

	crudService := &service.CrudService{}
	handler.NewMsgHandler(grpcServer, *crudService)

	grpcServer.Serve(lis)
}
