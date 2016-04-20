package main

import (
	"flag"
	// "io"
	"log"
	"net"
	"time"
)

var server = flag.String("listen", ":58080", "address server listens to")
var wait = flag.Int("wait", 10, "server wait time before closing")

func main() {
	flag.Parse()

	lsockFrom, err := net.Listen("tcp", *server)
	if err != nil {
		log.Fatalf("listen FROM error: %s\n", err.Error())
	}

	fromSock, err := lsockFrom.Accept()
	if err != nil {
		log.Printf("accept %s error: %s\n", *server, err.Error())
		return
	}

	buf := make([]byte, 128)
	n, err := fromSock.Read(buf)
	if err != nil {
		log.Printf("read error: %v\n", err)
		return
	}
	log.Printf("recv: %v\n", string(buf[:n]))

	fromSock.Write(buf[:n])
	log.Printf("response sent\n")

	time.Sleep(time.Duration(*wait) * time.Second)

	fromSock.Close()
}
