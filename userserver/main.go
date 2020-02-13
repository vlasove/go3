package main

import (
	"fmt"
	"log"

	"github.com/micro/go-micro"
	pb "github.com/vlasove/userserver/proto/user"
)

func main() {
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to database POSTGRES")
	}
	db.AutoMigrate(&pb.User{})
	repo := &UserRepository{db}

	srv := micro.NewService(
		micro.Name("user"),
		micro.Version("latest"),
	)
	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
