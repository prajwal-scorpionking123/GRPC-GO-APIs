package main

import (
	pb "GRPC-GO-LEARNING/greet/proto"
	"fmt"
	"log"
)
func (s *Server)GreetMany(in *pb.GreetRequest,stream pb.GreetService_GreetManyServer) error{
  log.Println("GreetMany Invoked with:",in)
  for i:=0;i<10;i++{
	res:=fmt.Sprintf("Hello %s",in.FirstName)
	stream.Send(&pb.GreetResponse{
		Result: res,
	})
  }
  return nil
}