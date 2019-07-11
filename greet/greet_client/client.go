package main

import (
	"log"

	"github.com/wendysanarwanto/go-grpc-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	// Create connection to gRPC server
	serverURL := "localhost:50051"
	grpcOption := grpc.WithInsecure() // We don't want to be bothered with setup SSL/TLS at the moment
	connection, err := grpc.Dial(serverURL, grpcOption)
	if err != nil {
		log.Fatalf("[ERROR] Failed connecting to gRPC server: %v", err)
		return
	}
	defer connection.Close() // Execute this line right before exiting this main function

	// Create the client
	client := greetpb.NewGreetServiceClient(connection)
	log.Printf("[INFO] Created client: %f", client)

}
