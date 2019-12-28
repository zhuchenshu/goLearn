package main

import (
	"net/rpc"
	"net"
	"github.com/lunny/log"
	"main/server"
	"net/rpc/jsonrpc"
)

func main()  {
	rpc.RegisterName("HelloService", new(server.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}