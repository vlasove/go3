package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"context"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
	pb "github.com/vlasove/shippyclient/proto/consignment"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

// nicePrint ...
func nicePrint(c *pb.Consignment) {
	log.Println("==================CONSIGNMENT DATA===============")
	log.Printf("Description: %s\n", c.GetDescription())
	log.Printf("Weight: %d\n", c.GetWeight())
	log.Printf("Containers: %v\n", c.GetContainers())
	log.Printf("VesselID: %s\n", c.GetVesselId())
	log.Println("=================================================")
	log.Println()
}

func main() {

	cmd.Init()

	client := pb.NewShippingServiceClient("shippyserver", microclient.DefaultClient)

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.TODO(), consignment)
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		nicePrint(v)
	}
}
