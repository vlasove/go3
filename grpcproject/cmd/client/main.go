package main

import (
	"context"
	"flag"
	"log"
	"strconv"

	"github.com/vlasove/grpcproject/pkg/api"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()

	arg, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	connection, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := api.NewSquarerClient(connection)
	res, err := client.Square(context.Background(), &api.SquareRequest{Arg: int32(arg)})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.GetAnswer())
}
