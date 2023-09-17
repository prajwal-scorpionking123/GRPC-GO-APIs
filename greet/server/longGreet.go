package main

import (
	pb "GRPC-GO-LEARNING/greet/proto"
	"fmt"
	"io"
	"log"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error{
   fmt.Println("SERVER LONG GREET INVOKED")
   res:=""
   for{
	 req,err:=stream.Recv()
	 if err==io.EOF{
		return stream.SendAndClose(&pb.GreetResponse{
			Result: res,
		})
	 }
	 if err!=nil{
		log.Fatalf("Error while reading client stream: %v\n", err)
	 }
	 res += "Hello " + req.FirstName + "!\n"
   }
}

	