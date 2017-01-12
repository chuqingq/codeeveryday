package main

import (
	"flag"
	"log"
	"net"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	address1 := flag.String("addr1", "117.78.4.30:20022", "address 1")
	address2 := flag.String("addr2", "127.0.0.1:22", "address 2")
	flag.Parse()

	var sock1 net.Conn
	var sock2 net.Conn
	var err error

	for {
		time.Sleep(1 * time.Second)

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
}

func pipe(sock1, sock2 *net.Conn) {
	buf := make([]byte, 64)
	for {
		n, err := (*sock1).Read(buf)
		if err != nil {
			log.Printf("sock[%v] read error: %v\n", *sock1, err)
			(*sock1).Close()
			*sock1 = nil
			return
		}

		if *sock2 == nil {
			continue
		}

		n, err = (*sock2).Write(buf[:n])
		if err != nil {
			log.Printf("sock[%v] write error: %v\n", *sock2, err)
			(*sock2).Close()
			*sock2 = nil
		}
	}
}
