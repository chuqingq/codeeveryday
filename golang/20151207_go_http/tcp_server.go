package main

import (
	"flag"
	// "io"
	// "fmt"
	"log"
	"net"
	"time"
)

var server = flag.String("listen", ":58080", "address server listens to")

// var wait = flag.Int("wait", 5, "wait seconds")
// var step = flag.Int("step", 0, "wait seconds")

func main() {
	flag.Parse()

	lsockFrom, err := net.Listen("tcp", *server)
	if err != nil {
		log.Fatalf("listen FROM error: %s\n", err.Error())
	}

	for true {
		fromSock, err := lsockFrom.Accept()
		if err != nil {
			log.Printf("accept %s error: %s\n", *server, err.Error())
			return
		}
		log.Printf("accept new socket\n")

		go handle(fromSock)
	}
}

func handle(fromSock net.Conn) {
	// fromSock.SetDeadline(2 * time.Second)
	buf := make([]byte, 128)
	// buf := []byte("123456")
	for true {
		// read
		n, err := fromSock.Read(buf)
		if err != nil {
			log.Printf("Read error: %v\n", err)
			return
		}
		log.Printf("Read: %v\n", string(buf[:n]))

		// write
		fromSock.SetWriteDeadline(time.Now().Add(3 * time.Second))

		n, err = fromSock.Write(buf[:n])
		if err != nil {
			log.Printf("Write error: %v\n", err)
			return
		}
		log.Printf("Write: %v\n", string(buf[:n]))

		// fromSock.Write(buf[:n])
		// log.Printf("response sent\n")

		// time.Sleep(time.Duration(*wait) * time.Second)
		// *wait += *step
	}

	fromSock.Close()
}
