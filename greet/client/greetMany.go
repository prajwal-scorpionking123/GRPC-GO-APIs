package main

import (
	pb "GRPC-GO-LEARNING/greet/proto"
	"context"
	"fmt"
	"io"
	"log"
)
func doGreetManyTime(client pb.GreetServiceClient){
   log.Println("SERVER INVOKED")
   stream,err:= client.GreetMany(context.Background(),&pb.GreetRequest{
	  FirstName: "PRAJWALP",
   })
   if err!=nil{
	 log.Fatal("error while getting res:",err)
   }
   for{
	res,err:=stream.Recv()
	if err==io.EOF{
		break
	}
	if err!=nil{
		log.Println("failed to stream",err)
	}
	fmt.Println("result:",res.Result)
   }

}