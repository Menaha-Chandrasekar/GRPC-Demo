package main

import (
	"context"
	"fmt"
	hw "grpc/helloworld"
	"net"

	"google.golang.org/grpc"
)

type server struct{
	hw.UnimplementedGreeterServer
}

func(s* server)SayHello(ctx context.Context,req *hw.HelloRequest)(*hw.HelloResponse,error){
	return &hw.HelloResponse{
		Message:fmt.Sprintf("Hello,%s! %d",req.Name,req.Age),
	},nil
}

func main(){
	lis,err:=net.Listen("tcp",":50051")
	if err!=nil{
		fmt.Printf("Failed to listen:%v",err)
		return
	}
	s:=grpc.NewServer()
	hw.RegisterGreeterServer(s,&server{})

	fmt.Println("Server listening on: 50051")
	if err:=s.Serve(lis);err !=nil{
		fmt.Printf("Failed to serve: %v",err)
	}
}