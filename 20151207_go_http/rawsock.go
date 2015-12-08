package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.ListenPacket("ip4:udp", "0.0.0.0")
	if err != nil {
		log.Fatalf("ListenPacket error: %v\n", err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	count := 0
	const MAXCOUNT = 200000
	before := time.Now()
	for {
		_, _, err := conn.ReadFrom(buf)
		if err != nil {
			log.Printf("ReadFrom error: %v\n", err)
		}

		// if addr.String() != "192.168.23.72" {
		// 	continue
		// }

		// log.Printf("ReadFrom %v\n", addr)
		count += 1
		if count == MAXCOUNT {
			log.Printf("recv: %f\n", float64(count)/time.Since(before).Seconds())
			count = 0
			before = time.Now()
		}
	}
}
