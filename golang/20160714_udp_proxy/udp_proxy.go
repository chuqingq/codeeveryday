package main

import (
	"flag"
	// "io"
	"log"
	"net"
)

var from = flag.String("from", ":20022", "the address this proxy will expose")
var to = flag.String("to", "127.0.0.1:22", "the address this proxy will connect to")

func main() {
	flag.Parse()

	fromAddr, err := net.ResolveUDPAddr("udp", *from)
	if err != nil {
		log.Fatalf("from is invalid: %s\n", err.Error())
	}

	fromSock, err := net.ListenUDP("udp", fromAddr)
	if err != nil {
		log.Fatalf("ListenUDP error: %s\n", err.Error())
	}

	toAddr, err := net.ResolveUDPAddr("udp", *to)
	if err != nil {
		log.Fatalf("to is invalid: %s\n", err.Error())
	}

	toSock, err := net.DialUDP("udp", nil, toAddr)
	if err != nil {
		log.Fatalf("DialUDP error: %s\n", err.Error())
	}

	b := make([]byte, 1024)
	for {
		n, err := fromSock.Read(b)
		if err != nil {
			log.Fatalf("Read error: %v\n", err)
		}

		log.Printf("read len: %d\n", n)

		n2, err := toSock.Write(b[:n])
		if n2 != n {
			log.Fatalf("write n not match! n2=%d, n=%d\n", n2, n)
		}

		if err != nil {
			log.Fatalf("Write error: %v\n", err)
		}
	}
}
