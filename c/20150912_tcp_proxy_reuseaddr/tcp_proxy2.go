package main

import (
	"flag"
	// "io"
	"log"
	"net"
	"time"
)

func main() {
	var address1 = flag.String("address1", ":20022", "the address1 this proxy listen to")
	var address2 = flag.String("address2", ":20023", "the address2 this proxy listen to")
	flag.Parse()

	lsock1, err := net.Listen("tcp", *address1)
	if err != nil {
		log.Fatalf("listen address1(%s) error: %s\n", *address1, err.Error())
	}

	lsock2, err := net.Listen("tcp", *address2)
	if err != nil {
		log.Fatalf("listen address2(%s) error: %s\n", *address2, err.Error())
	}

	var sock1 net.Conn
	var sock2 net.Conn
	for {
		time.Sleep(1 * time.Second)

		if sock1 == nil {
			sock1, err = lsock1.Accept()
			if err != nil {
				log.Fatalf("accept address1(%s) error: %s\n", *address1, err.Error())
			}
			log.Printf("sock1: %v\n", sock1)

			go pipe(&sock1, &sock2)
		}

		if sock2 == nil {
			sock2, err = lsock2.Accept()
			if err != nil {
				log.Fatalf("accept address2(%s) error: %s\n", *address2, err.Error())
			}
			log.Printf("sock2: %v\n", sock2)

			go pipe(&sock2, &sock1)
		}
	}
}

func pipe(sock1, sock2 *net.Conn) {
	buf := make([]byte, 64)
	for {
		n, err := (*sock1).Read(buf)
		if err != nil {
			log.Printf("sock[%v] closed\n", *sock1)
			(*sock1).Close()
			*sock1 = nil
			return
		}

		if *sock2 == nil {
			continue
		}

		n, err = (*sock2).Write(buf[:n])
		if err != nil {
			(*sock2).Close()
			*sock2 = nil
		}
	}
	// _, err := io.Copy(*sock2, *sock1)
	// log.Printf("copy[%v->%v] error: %v\n", *sock1, *sock2, err)
	// if err == nil {

	// } else {
	// 	(*sock2).Close()
	// 	*sock2 = nil
	// }
	// time.Sleep(10 * time.Second)
}
