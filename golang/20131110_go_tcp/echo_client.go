package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	const count = 10000
	buf := make([]byte, 1024)
	msg := []byte("hello world")

	start := time.Now()
	for i := 0; i < count; i++ {
		//准备要发送的字符串
		_, err := conn.Write(msg)
		if err != nil {
			println("Write Buffer Error:", err.Error())
			break
		}
		// fmt.Printf("write: %s\n", msg)

		//从服务器端收字符串
		n, err := conn.Read(buf)
		if err != nil {
			println("Read Buffer Error:", err.Error())
			break
		}
		// fmt.Printf("read: %s\n", string(buf[0:n]))
		if n != len(msg) {
			log.Printf("length not match")
		}

		// //等一秒钟
		// time.Sleep(time.Second)
	}

	log.Printf("%v", time.Now().Sub(start)/count)
}
