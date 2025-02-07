package main

import (
	"log"
	"net"

	greetpb "github.com/mullayam/greet/proto"
	"google.golang.org/grpc"
)

var addr string = ":50051"

type Server struct {
	greetpb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("TCP Server listening at %v", lis.Addr())
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
