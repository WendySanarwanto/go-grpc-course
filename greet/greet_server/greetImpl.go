package main

import ( 
	"context"
	"log"
	"strconv"
	"time"

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

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	log.Printf("[INFO] GreetManyTimes function was invoked with %v", req)
	// Send greet 10 times
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	numGretings := int(req.GetNumGreetings())

	for i:=1; i<=numGretings; i++ {
		result := "Hello, "+firstName+" "+lastName+". Number: "+ strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(200 * time.Millisecond)
	}
	return nil
}