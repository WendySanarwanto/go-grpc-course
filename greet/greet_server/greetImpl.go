package main

import ( 
	"context"
	"log"
	"github.com/wendysanarwanto/go-grpc-course/greet/greetpb"	
)

// Implementation of Greet method
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("[INFO] Greet function was invoked with %v", req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetFirstName()
	result := "Hello, " + firstName + " " + lastName
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}
