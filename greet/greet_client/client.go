package main

import (
	"context"
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

	// Invoke Greet function
	doUnaryRequest(client)
}

func doUnaryRequest(client greetpb.GreetServiceClient) {
	log.Println("Starting to do a Unary RPC...")
	// Invoke Greet function
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ada",
			LastName: "Wong",
		},
	}
	res, err := client.Greet(context.Background(), req )
	if err != nil {
		log.Fatalf("[ERROR] Got error when called 'Greet' RPC: %v", err)
	}

	log.Printf("Response from greet: %v",res.Result )	
}
