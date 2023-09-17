package main

import (
	pb "GRPC-GO-LEARNING/greet/proto"
	"context"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error){
  log.Println("GREET FUNCTION WAS INVOKED WITH:",in)
  return &pb.GreetResponse{
	Result: "hello "+in.FirstName,
  },nil
}