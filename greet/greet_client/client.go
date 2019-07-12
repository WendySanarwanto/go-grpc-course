package main

import (
	"context"
	"io"
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
	firstName := "Ada"
	lastName := "Wong"
	numGreetings := 10
	readStreamedChunkDelay := 1000
	doUnaryRequest(firstName, lastName, client)
	doServerStreaming(firstName, lastName, int32(numGreetings), int32(readStreamedChunkDelay), client)
}

func doUnaryRequest(firstName string, lastName string, client greetpb.GreetServiceClient) {
	log.Println("[INFO] Starting to do a Unary RPC...")
	// Invoke Greet function
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: firstName,
			LastName: lastName,
		},
	}
	res, err := client.Greet(context.Background(), req )
	if err != nil {
		log.Fatalf("[ERROR] Got error when called 'Greet' RPC: %v", err)
	}

	log.Printf("[INFO] Response from greet: %v",res.Result )	
}

func doServerStreaming(firstName string, lastName string, numGreetings int32, 
											readStreamedChunkDelay int32, client greetpb.GreetServiceClient) {
	log.Println("[INFO] Starting to do a Server Streaming RPC...")
	// Invoke GreetManyTimes function
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: firstName,
			LastName: lastName,
		},
		NumGreetings: numGreetings,
		ReadStreamedChunkDelay: readStreamedChunkDelay,	
	}

	resStream, err := client.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf(`[ERROR] Invoking 'GreetManyTimes' is failed: %v`, err);
		return;
	}

	// Read response items from the stream until reacthed End of Stream
	for {
		msg, err := resStream.Recv()
		// Identify if the error indicates End of stream or something else
		if err == io.EOF {
			// we've reacthed the end of the stream
			break;
		} else if err != nil {
			log.Fatalf(`[ERROR] Error while reading response stream: %v`, err)
		}
		log.Printf(`[INFO] Response from GreetManyTimes: %v`, msg.GetResult())
	}
}
