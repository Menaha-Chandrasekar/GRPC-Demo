package main

import (
"context"
"fmt"
"net"
"sync"


pb "grpc/task"

// "go.mongodb.org/mongo-driver/bson/primitive"
"go.mongodb.org/mongo-driver/bson/primitive"
"google.golang.org/grpc"
)


type taskServiceserver struct{
mu sync.Mutex
tasks map[string]*pb.Task
pb.UnimplementedTaskServiceServer
}

func (s*taskServiceserver)AddTask(ctx context.Context,req*pb.Task)(*pb.TaskResponse,error){
s.mu.Lock()
defer s.mu.Unlock()
taskID:=generateID()
req.Id =taskID
s.tasks[taskID]=req

return &pb.TaskResponse{
Id:taskID},nil
}

func(s*taskServiceserver)GetTasks(ctx context.Context,req *pb.Empty)(*pb.TaskList,error){
s.mu.Lock()
defer s.mu.Unlock()

tasks:=make([]*pb.Task,0,len(s.tasks))
for _,task:=range s.tasks{
tasks =append(tasks, task)
}
return &pb.TaskList{Tasks: tasks},nil
}

func generateID() string{

id:=primitive.NewObjectID()
taskID:=id.Hex()

return taskID
}
func main(){
lis,err:=net.Listen("tcp", ":50051")
if err!=nil{
fmt.Printf("Failed to lIsten:%v", err)
return
}
s:=grpc.NewServer()
pb.RegisterTaskServiceServer(s,&taskServiceserver{
tasks: make(map[string]*pb.Task),
})
fmt.Println("Server listening on:50051")
if err:=s.Serve(lis);err!=nil{
fmt.Printf("failed to server:%v",err)
}
}