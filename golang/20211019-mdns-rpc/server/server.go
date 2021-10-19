package main

import (
	"log"
	mrpc "mdns_rpc_sample"
	"mdns_rpc_sample/message"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	arith := new(message.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":8081")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
	mrpc.RegisterService("_foobar._tcp", 8081)
	select {}
}
