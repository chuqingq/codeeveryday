package main

import (
	"log"
	mrpc "mdns_rpc_sample"
	"mdns_rpc_sample/message"
)

func main() {
	{
		args := &message.Args{7, 8}
		var reply int
		err := mrpc.RPCCall("_foobar._tcp", "Arith.Multiply", args, &reply)
		if err != nil {
			log.Printf("RPCCall error: %v", err)
			return
		}
		log.Printf("Arith.Multiply: %v %v %v", args.A, args.B, reply)
	}

	// 再来一次
	{
		args := &message.Args{7, 8}
		var reply int
		err := mrpc.RPCCall("_foobar._tcp", "Arith.Multiply", args, &reply)
		if err != nil {
			log.Printf("RPCCall error: %v", err)
			return
		}
		log.Printf("Arith.Multiply: %v %v %v", args.A, args.B, reply)
	}
}
