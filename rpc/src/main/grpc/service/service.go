package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	grpc2 "main/grpc/api"
)

const PORT  = ":50051"

type server struct {}

func (s *server)SayHello(ctx context.Context, in *grpc2.HelloRequest)(*grpc2.HelloReply,error){
	return &grpc2.HelloReply{Message:"hello"},nil
}

func main(){
	//监听端口
	lis,err := net.Listen("tcp",PORT)
	if err != nil{
		return
	}
	//创建一个grpc 服务器
	s := grpc.NewServer()
	//注册事件
	grpc2.RegisterGreeterServer(s,&server{})
	//处理链接
	s.Serve(lis)
}