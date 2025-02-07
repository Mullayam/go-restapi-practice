package main

import (
	"context"
	"log"

	greetpb "github.com/mullayam/greet/proto"
)

func (s *Server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", req)
	return &greetpb.GreetResponse{
		Message: "Hello " + req.FirstName,
	}, nil
}
