package main

import (
	"log"

	greetpb "github.com/mullayam/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, _err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if _err != nil {
		log.Fatalf("could not connect: %v", _err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	doGreet(c)
}
