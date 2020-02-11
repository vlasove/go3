package main

import (
	"log"
	"net"

	"github.com/vlasove/grpcproject/pkg/api"
	"github.com/vlasove/grpcproject/pkg/squarer"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	microService := &squarer.GRPCServer{}
	api.RegisterSquarerServer(server, microService)

	Listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Serve(Listener); err != nil {
		log.Fatal(err)
	}

}
