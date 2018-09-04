// consignment-cli/cli.go
package main

import (
	pb "shippo/consignment-service/proto/consignment"
	"io/ioutil"
	"encoding/json"
	"google.golang.org/grpc"
	"log"
	"os"
	"context"
)

const (
address         = "localhost:50051"
defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignemt *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignemt)
	return consignemt, nil
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewShippingServiceClient(conn)

	// Contact the server and print out its response.
	file := defaultFilename
	log.Println(os.Args[0])
	if len(os.Args) > 1{
		file = os.Args[1]
	}
	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

}