package main

import (
	"flag"
	"io"
	"log"
	"net"
)

var address1 = flag.String("address1", ":20022", "the address1 this proxy listen to")
var address2 = flag.String("address2", ":20023", "the address2 this proxy listen to")

func main() {
	flag.Parse()

	lsock1, err := net.Listen("tcp", *address1)
	if err != nil {
		log.Fatalf("listen address1(%s) error: %s\n", *address1, err.Error())
	}

	lsock2, err := net.Listen("tcp", *address2)
	if err != nil {
		log.Fatalf("listen address2(%s) error: %s\n", *address2, err.Error())
	}

	for {
		sock1, err := lsock1.Accept()
		if err != nil {
			log.Fatalf("accept address1(%s) error: %s\n", *address1, err.Error())
		}

		sock2, err := lsock2.Accept()
		if err != nil {
			log.Fatalf("accept address2(%s) error: %s\n", *address2, err.Error())
		}
		
		go pipe(sock1, sock2)
		go pipe(sock2, sock1)
	}
}

func pipe(sock1, sock2 net.Conn) {
	io.Copy(sock1, sock2)
	sock1.Close()
}
