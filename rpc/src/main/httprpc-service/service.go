package main

import (
	"net/rpc"
	"io"
	"net/http"
	"main/server"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.RegisterName("HelloService", new(server.HelloService))

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	for {
		http.ListenAndServe(":1234", nil)
	}
}