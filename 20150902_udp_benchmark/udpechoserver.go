package main

import (
	// "flag"
	// "fmt"
	"log"
	"net"
	// "time"
	// "runtime"
)

func main() {
	localaddr := "0.0.0.0:20000"

	local_addr, err := net.ResolveUDPAddr("udp", localaddr)
	if err != nil {
		log.Fatalf("localaddr resolve error: %v\n", err)
	}

	conn, err := net.ListenUDP("udp", local_addr)
	if err != nil {
		log.Fatalf("ListenUDP error: %v\n", err)
	}

	buf := make([]byte, 1024*1024)
	for {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Printf("read error: %v\n", err)
		}
		log.Printf("recv %s\n", string(buf[0:n]))

		// time.Sleep(2 * time.Second)

		_, err = conn.WriteTo(buf[0:n], addr)
		if err != nil {
			log.Printf("write error: %v\n", err)
			continue
		}
	}
}
