package main

import (
	pb "GRPC-GO-LEARNING/greet/proto"
	"io"
	"log"
)
func (s *Server)GreetEveryOne(stream pb.GreetService_GreetEveryOneServer)error{
  log.Println("GreetEveryOne INVOKED with:")
  for{
	req,err:=stream.Recv()
	if err==io.EOF{
		return nil
	}
	if err!=nil{
		log.Fatal("failed to recieve the data",err)
	}
	res:="Hello "+req.FirstName
	err = stream.Send(&pb.GreetResponse{
		Result: res,
	})
	if err!=nil{
		log.Fatal("failed to send the data",err)
	}
  }
}