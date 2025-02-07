package main

import (
	"context"
	"log"

	greetpb "github.com/mullayam/greet/proto"
)

func doGreet(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		FirstName: "Mullayam",
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Message)
}
