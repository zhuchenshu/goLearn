package main

import (
	"net/rpc"
	"net"
	"github.com/lunny/log"
	"main/server"
)

func main()  {
	rpc.RegisterName("HelloService", new(server.HelloService))

	listener , err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen tcp err:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("accept err:", err)
	}

	rpc.ServeConn(conn)
}