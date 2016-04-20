package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

var server = flag.String("server", ":58080", "server address to connect to")
var wait = flag.Int("wait", 10, "client wait time before exiting")

func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", *server)
	if err != nil {
		log.Printf("dial error: %v\n", err)
		return
	}

	fmt.Fprintf(conn, "GET / HTTP/1.1\r\n\r\n")

	data := make([]byte, 128)
	n, err := conn.Read(data)
	if err != nil {
		log.Printf("read error: %v\n", err)
		return
	}
	log.Printf("read: %v\n", string(data[:n]))

	time.Sleep(time.Duration(*wait) * time.Second)

	fmt.Fprintf(conn, "GET /2 HTTP/1.1\r\n\r\n")

	n, err = conn.Read(data)
	if err != nil {
		log.Printf("read error: %v\n", err)
		return
	}
	log.Printf("read: %v\n", string(data[:n]))
}
