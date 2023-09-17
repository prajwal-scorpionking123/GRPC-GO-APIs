package main

import (
	pb "GRPC-GO-LEARNING/greet/proto"
	"context"
	"log"
)
func doGreet(client pb.GreetServiceClient){
   log.Println("CLIENT IS CALLING SERVER with ",client)

   response,err:= client.Greet(context.Background(),&pb.GreetRequest{
	FirstName: "PRAJWALP",
   })
   if err!=nil{
	 log.Fatal("error which calling server::",err)
   }
   log.Println(response.Result)
}