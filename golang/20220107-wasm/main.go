package main

import (
	"fmt"
	"log"
	"net"
	"syscall/js"
)

func main() {
	log.Printf("123")
	alert := js.Global().Get("alert")
	alert.Invoke("Hello World!")
	test()
}

func test() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Printf("dial error: %s\n", err.Error())
		return
	}

	fmt.Fprintf(conn, "GET / HTTP/1.1\r\n\r\n")

	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil {
		fmt.Printf("read error: %#v\n", err)
		return
	}

	fmt.Printf("read: %s\n", data[:n])
}
