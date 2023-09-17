package main

import (
	pb "GRPC-GO-LEARNING/greet/proto"
	"context"
	"io"
	"log"
	"time"
)
func doGreetEveryOne(client pb.GreetServiceClient){
  log.Println("GREETING EVERYONE!")
  stream,err:=client.GreetEveryOne(context.Background())
  if err!=nil{
	log.Fatal("FAILED TO CALL GREET EVERYONE")
  }
  reqs := []*pb.GreetRequest{
	{FirstName: "PRAJWAL"},
	{FirstName: "RAJ"},
	{FirstName: "NAKUL"},
  }
  waitc:=make(chan struct{})
  go func() {
    for _,req:=range reqs{
	    err:=stream.Send(&pb.GreetRequest{
			FirstName: req.FirstName,
		})
		if err!=nil{
			log.Fatal("FAILED TO Send req")
		}
		log.Println("SEND Request:",req.FirstName)
		time.Sleep(1 * time.Second)
	}
	stream.CloseSend()
  }()
  go func() {
    for{
	 res,err:=stream.Recv()
     if err==io.EOF{
		break
	 }
	 if err!=nil{
		log.Fatalln("FAILED RECIEVE DATA")
	 }
	 log.Printf("Received: %v\n", res.Result)
	}
	close(waitc)
  }()
  <-waitc
}