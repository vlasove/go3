package main

import (
	fmt "fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	p1 := &Person{
		Name:     "Bob",
		Lastname: "Dylan",
		Age:      int64(33),
	}

	data, err := proto.Marshal(p1)
	if err != nil {
		log.Fatal("during marshalling happend error: ", err)
	}
	fmt.Println(data)

	p2 := &Person{}
	err = proto.Unmarshal(data, p2)
	if err != nil {
		log.Fatal("during unmarshalling happend error: ", err)
	}

	fmt.Println("Name:", p2.GetName())
	fmt.Println("Lastname:", p2.GetLastname())
	fmt.Println("Age:", p2.GetAge())
}
