package main

import (
	"net/rpc"
	"github.com/lunny/log"
	"fmt"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dial err:", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "rpc", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(reply)
}