package main

import (
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro"
	pb "github.com/vlasove/shippyserver/proto/consignment"
	vesselProto "github.com/vlasove/shippyserver/proto/vessel"
)

const (
	defaultHost = "localhost:27017"
)

func main() {

	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)

	defer session.Close()

	if err != nil {

		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	srv := micro.NewService(

		micro.Name("shippyserver"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("shippyvessel", srv.Client())

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
