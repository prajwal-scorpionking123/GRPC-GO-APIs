package main

import (
	"log"

	pb "GRPC-GO-LEARNING/greet/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:5051"

type Client struct{
	
}
func main(){
  conn,err:=grpc.Dial(addr,grpc.WithTransportCredentials(insecure.NewCredentials()))
  if err!=nil{
	log.Fatal("failed to connect to server",err)
  }
  client:=pb.NewGreetServiceClient(conn)
  doGreetEveryOne(client)
  defer conn.Close()
}