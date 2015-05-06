package main

import (
	"flag"
	"io"
	"log"
	"net"
)

var from = flag.String("from", ":20022", "the address this proxy will expose")
var to = flag.String("to", "127.0.0.1:22", "the address this proxy will connect to")

func main() {
	flag.Parse()

	lsockFrom, err := net.Listen("tcp", *from)
	if err != nil {
		log.Fatalf("listen FROM error: %s\n", err.Error())
	}

	for {
		fromSock, err := lsockFrom.Accept()
		if err != nil {
			log.Fatalf("accept FROM(%s) error: %s\n", *from, err.Error())
		}

		toSock, err := net.Dial("tcp", *to)
		if err != nil {
			log.Fatalf("Dial(%s) error: %s\n", *to, err.Error())
			continue
		}

		go io.Copy(fromSock, toSock)
		go io.Copy(toSock, fromSock)
	}
}
