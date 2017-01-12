package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	go udpServer()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := net.Dial("udp", "127.0.0.1:3001")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			log.Printf("http dialudp error: %v\n", err.Error())
			return
		}
		defer conn.Close()

		buffer := []byte("hello world")
		conn.Write(buffer)
		// buffer := make([]byte, 64)
		len, err := conn.Read(buffer)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			log.Printf("http udp read error: %v\n", err.Error())
			return
		}
		if string(buffer) != "hello world" {
			log.Printf("udp msg not match: %v\n", string(buffer))
		}
		w.Write(buffer[:len])
	})
	log.Fatalf("ListenAndServe: %v\n", http.ListenAndServe(":3000", nil))
}

func udpServer() {
	laddr, err := net.ResolveUDPAddr("udp", ":3001")
	if err != nil {
		log.Fatalf("ResolveUDPAddr error: %v\n", err)
	}

	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		log.Fatalf("listen udp error: %v\n", err.Error())
	}

	b := make([]byte, 64)
	for {
		len, addr, err := conn.ReadFrom(b)
		if err != nil {
			log.Printf("readfrom error: %v\n", err.Error())
			continue
		}
		conn.WriteTo(b[:len], addr)
	}
}
