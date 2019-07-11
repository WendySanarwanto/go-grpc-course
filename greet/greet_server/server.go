package main

import ( 
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/wendysanarwanto/go-grpc-course/greet/greetpb"
)

type server struct{} // Declared an empty struct, named as `server`

func main() {
	serverURL := "0.0.0.0:50051"
	log.Printf("[INFO] Starting the gRPC server to listen at %v ...", serverURL)

	listener, err := net.Listen("tcp", serverURL)
	if err != nil {
		log.Fatalf("[ERROR] Failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	// Create an instance of `server` and 
	// pass it as 2nd arg of RegisterGreetServiceServer -> Why ?
	greetpb.RegisterGreetServiceServer(s, &server{});

	if err := s.Serve(listener); err != nil {
		log.Fatalf("[ERROR] Failed to serve: %v", err)
		return
	}	
}