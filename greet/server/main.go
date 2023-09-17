package main

import (
	pb "GRPC-GO-LEARNING/greet/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
}

var addr string = "0.0.0.0:5051"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("failed to list on:", err)
	}
	log.Println("Listening on:", addr)
	server := grpc.NewServer()
	pb.RegisterGreetServiceServer(server,&Server{})
	if err := server.Serve(lis); err != nil {
		log.Fatal("failed to serve", err)
	}
}
