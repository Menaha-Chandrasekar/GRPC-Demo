package main

import (
	"context"
	"fmt"
	hw "grpc/customer"
	"net"

	"google.golang.org/grpc"
)

type server1 struct{
	hw.UnimplementedCustomerServiceServer
}

func(s* server1)GetCustomer(ctx context.Context,req *hw.Task)(*hw.TaskResponse,error){
	return &hw.TaskResponse{
		Message:fmt.Sprintf("Hello,%s! %d %d %d ",req.Name,req.Accountid,req.Balance,req.Bankid),
	},nil
}

func main(){
	lis,err:=net.Listen("tcp",":50051")
	if err!=nil{
		fmt.Printf("Failed to listen:%v",err)
		return
	}
	s:=grpc.NewServer()
	hw.RegisterCustomerServiceServer(s,&server1{})

	fmt.Println("Server listening on: 50051")
	if err:=s.Serve(lis);err !=nil{
		fmt.Printf("Failed to serve: %v",err)
	}
}