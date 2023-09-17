package main

import (
	pb "GRPC-GO-LEARNING/greet/proto"
	"context"
	"log"
	"time"
)

func doLongGreet(client pb.GreetServiceClient) {
	log.Println("SERVER INVOKED")
	reqs := []*pb.GreetRequest{
		{FirstName: "PRAJWAL"},
		{FirstName: "RAJ"},
		{FirstName: "NAKUL"},
	}
	stream, err := client.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
