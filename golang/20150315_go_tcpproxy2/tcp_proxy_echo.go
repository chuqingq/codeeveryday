package main

/*
主动连接到address，并作为echo_server
*/

import (
	"flag"
	"io"
	"log"
	"net"
)

var address = flag.String("address", "127.0.0.1:20022", "the address this proxy will expose")

func main() {
	flag.Parse()

	sock, err := net.Dial("tcp", *address)
	if err != nil {
		log.Fatalf("Dial(%s) error: %s\n", *address, err.Error())
		return
	}

	io.Copy(sock, sock)
}
