package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"runtime"
)

func main() {
	num_cores := runtime.NumCPU()
	runtime.GOMAXPROCS(num_cores)

	localaddr := flag.String("local", "", "local addr")
	addr := flag.String("server", "127.0.0.1:20000", "server")
	c := flag.Int("c", num_cores, "connections")
	recv := flag.Bool("recv", true, "recv")
	flag.Parse()

	local_addr, err := net.ResolveUDPAddr("udp", *localaddr)
	if err != nil {
		log.Fatalf("localaddr resolve error: %v\n", err)
	}

	server_addr, err := net.ResolveUDPAddr("udp", *addr)
	if err != nil {
		log.Fatalf("addr resolve error: %v\n", err)
	}

	log.Printf("cores: %d, connections: %d, recv: %v, server: %s\n", num_cores, *c, *recv, *addr)

	for i := 0; i < *c; i++ {
		go func() {
			conn, err := net.ListenUDP("udp", local_addr)
			if err != nil {
				log.Fatalf("dialudp error: %v\n", err)
			}

			buf := make([]byte, 32)
			for {
				_, err = conn.WriteTo(buf, server_addr)
				if err != nil {
					log.Printf("write error: %v\n", err)
					continue
				}

				if *recv {
					_, err = conn.Read(buf)
					if err != nil {
						log.Printf("read error: %v\n", err)
					}
				}
			}
		}()
	}
	var wait string
	fmt.Scanf("%s", &wait)
}
