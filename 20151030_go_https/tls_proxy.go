package main

import (
	"flag"
	"io"
	"log"
	//"net"
	"crypto/tls"
)

var from = flag.String("from", ":443", "the address this proxy will expose")
var to = flag.String("to", "122.11.38.57:443", "the address this proxy will connect to")

func main() {
	flag.Parse()

	cert, err := tls.LoadX509KeyPair(
		"./cert.pem",
		"./key.pem",
	)
	if err != nil {
		log.Fatalf("load cert error: %v\n", err)
	}

	config := tls.Config{Certificates: []tls.Certificate{cert}}

	lsockFrom, err := tls.Listen("tcp", *from, &config)
	if err != nil {
		log.Fatalf("listen FROM error: %s\n", err.Error())
	}

	for {
		fromSock, err := lsockFrom.Accept()
		if err != nil {
			log.Fatalf("accept FROM(%s) error: %s\n", *from, err.Error())
		}

		toSock, err := tls.Dial("tcp", *to, &tls.Config{InsecureSkipVerify: true})
		if err != nil {
			log.Fatalf("Dial(%s) error: %s\n", *to, err.Error())
			continue
		}

		go io.Copy(fromSock, toSock)
		go io.Copy(toSock, fromSock)
	}
}
