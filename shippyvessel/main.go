package main

import (
	"log"
	"os"

	"github.com/micro/go-micro"
	pb "github.com/vlasove/shippyvessel/proto/vessel"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	vessels := []*pb.Vessel{
		{Id: "vessel1", Name: "Bob", MaxWeight: 2000000, Capacity: 1000},
	}
	var repo Repository
	repo.Create(vessels[0])
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}
	session, err := CreateSession(host)
	defer session.Close()
	if err != nil {
		log.Fatal("Error connecting to db mongo")
	}

	srv := micro.NewService(
		micro.Name("shippyvessel"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{session})

}
