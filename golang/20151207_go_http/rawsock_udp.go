package main

import (
	"log"
	"net"
	//"time"
)

func main() {
	conn, err := net.ListenPacket("ip4:udp", "0.0.0.0")
	if err != nil {
		log.Fatalf("ListenPacket error: %v\n", err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Printf("ReadFrom error: %v\n", err)
		}

		// if addr.String() != "192.168.23.72" {
		// 	continue
		// }

		if addr.String() != "127.0.0.1" {
			continue
			log.Printf("ReadFrom %v\n", addr)
		}
		log.Printf("ReadFrom %v, %v, %v\n", addr, n, string(buf[8:n]))

		// conn.WriteTo(buf[:n], addr)
	}
}

/*
$ sudo ./rawsock_udp 
2017/08/11 16:09:24 ReadFrom 127.0.0.1, 24, 123123123123123

$ nc -u 127.0.0.1 12345
123123123123123

*/
