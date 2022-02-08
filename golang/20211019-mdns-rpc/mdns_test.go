package mdns_rpc_sample

import (
	"errors"
	"mdns_rpc_sample/message"
	"testing"
)

// RPC protocol

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func TestRPCCall(t *testing.T) {
	// server
	rpcserver := NewRPC()
	rpcserver.RegisterService("_foobar._tcp", new(message.Arith))
	defer rpcserver.Close()

	// client
	{
		rpcclient := NewRPC()
		args := &Args{7, 8}
		var reply int
		err := rpcclient.Call("_foobar._tcp", "Arith.Multiply", args, &reply)
		if err != nil {
			t.Fatalf("RPCCall error: %v", err)
		}
		if reply != 56 {
			t.Fatalf("Arith.Multiply error: %v*%v!=%v", args.A, args.B, reply)
		}
		rpcclient.Close()
	}

	// client again
	{
		rpcclient := NewRPC()
		args := &Args{6, 9}
		var reply int
		err := rpcclient.Call("_foobar._tcp", "Arith.Multiply", args, &reply)
		if err != nil {
			t.Fatalf("RPCCall error: %v", err)
			return
		}
		if reply != 54 {
			t.Fatalf("Arith.Multiply error: %v*%v!=%v", args.A, args.B, reply)
		}
	}
}
