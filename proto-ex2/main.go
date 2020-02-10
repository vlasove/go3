package main

import (
	fmt "fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	b1 := Bloger{
		Name:     "Bob",
		Lastname: "Dylan",
		Post: &Post{
			Likes:    int32(100),
			Dislikes: int32(5000),
		},
	}

	data, err := proto.Marshal(&b1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

}
