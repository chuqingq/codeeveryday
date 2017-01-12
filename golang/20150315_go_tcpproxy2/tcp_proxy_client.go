package main

/*
主动连接到address，并作为proxy的client
*/

import (
	"flag"
	// "io"
	"log"
	"net"
	"fmt"
	"time"
)

var address = flag.String("address", "127.0.0.1:20023", "the address this proxy will expose")

func main() {
	flag.Parse()

	sock, err := net.Dial("tcp", *address)
	if err != nil {
		log.Fatalf("Dial(%s) error: %s\n", *address, err.Error())
		return
	}
    defer sock.Close()
 
    buf := make([]byte, 1024)
 
    for i := 0; i < 5; i++ {
        //准备要发送的字符串
        msg := fmt.Sprintf("Hello World, %03d", i)
        n, err := sock.Write([]byte(msg))
        if err != nil {
            log.Printf("Write Buffer Error: %s\n", err.Error())
            break
        }
        log.Printf("write: %s\n", msg)
 
        //从服务器端收字符串
        n, err = sock.Read(buf)
        if err !=nil {
            log.Printf("Read Buffer Error: %s\n", err.Error())
            break
        }
        log.Printf("read: %s\n", string(buf[0:n]))
 
        //等一秒钟
        time.Sleep(time.Second)
    }
}
