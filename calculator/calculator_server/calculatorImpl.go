package main;

import (
	"context"
	"log"

	"github.com/wendysanarwanto/go-grpc-course/calculator/calculatorpb"
)

// Sum - Do add/Sum 2 numbers 
func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Printf("[INFO] Sum function was invoked with %v", req)
	left := req.Payload.Left
	right := req.Payload.Right

	total := left + right

	res := &calculatorpb.SumResponse{
		Total: total,
	}

	return res, nil
}
