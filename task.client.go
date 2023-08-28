package main

import (
"context"
"fmt"
"log"

// hw "Grpc-demo/helloworld"
pb "grpc/task"

"google.golang.org/grpc"
)

func main(){
conn,err:=grpc.Dial("localhost:50051",grpc.WithInsecure())
if err!=nil{
log.Fatal("failed1 %v",err)
}
defer conn.Close()
// client1:=hw.NewGreeterClient(conn)
client2:=pb.NewTaskServiceClient(conn)
// name:="thash"
// age:=18

task:=&pb.Task{
Title:"Buy groceries",
}

// response,err:=client1.SayHello(context.Background(),&hw.HelloRequest{Name:name,Age:int32(age)})
addres,err:=client2.AddTask(context.Background(),task)
if err!=nil{
log.Fatal("failed2 %v",err)
}

fmt.Printf("Response of add :%s\n",addres.Id)

taskres,err:=client2.GetTasks(context.Background(),&pb.Empty{})
if err!=nil{
log.Fatal("failed2 %v",err)
}
fmt.Printf("Response of get :")
for _,task:=range taskres.Tasks{
fmt.Printf("ID:%s,Title:%s,Completed:%v\n",task.Id,task.Title,task.Completed)
}
}