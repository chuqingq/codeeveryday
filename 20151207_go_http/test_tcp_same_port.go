package main

import (
	"flag"
	"log"
	"net"
	"runtime"
	"time"
)

func main() {
	address1 := flag.String("addr1", "192.168.13.165:10000", "address 1")
	address2 := flag.String("addr2", "127.0.0.1:10000", "address 2")
	flag.Parse()

	addr1, err := net.ResolveTCPAddr("tcp", *address1)
	if err != nil {
		log.Fatalf("resolve addr1 error: %v\n", err)
	}

	addr2, err := net.ResolveTCPAddr("tcp", *address2)
	if err != nil {
		log.Fatalf("resolve addr2 error: %v\n", err)
	}

	net.DialTCP()

	if sock1 == nil {
		sock1, err = net.Dial("tcp", *address1)
		if err != nil {
			log.Printf("dial address1[%s] error: %v\n", *address1, err)
			continue
		}
		log.Printf("dial address1[%s] success\n", *address1)

		go pipe(&sock1, &sock2)
	}

	if sock2 == nil {
		sock2, err = net.Dial("tcp", *address2)
		if err != nil {
			log.Printf("dial address2[%s] error: %v\n", *address2, err)
			continue
		}
		log.Printf("dial address2[%s] success\n", *address2)

		go pipe(&sock2, &sock1)
	}
}
