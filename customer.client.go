package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc/customer"

	"google.golang.org/grpc"
)

func main(){
	conn,err:=grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err !=nil{
		log.Fatalf("Failed to connect: %v",err)
	}
	defer conn.Close()
	client:=pb.NewCustomerServiceClient(conn)

	name:="menaha";
	var accountid int32
	accountid=908129;
	var balance int32
	balance=7000;
	var bankid int32
	bankid=36784;
	response,err:=client.GetCustomer(context.Background(),&pb.Task{Name:name,Accountid:accountid,Balance:balance,Bankid:bankid})
	if err!=nil{
		log.Fatalf("Failed to call Customer: %v",err)
	}
	fmt.Printf("Response:%s\n",response.Message)
}