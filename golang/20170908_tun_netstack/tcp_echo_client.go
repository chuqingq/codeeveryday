package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.109:8888")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	buf := []byte("hello world")

	count := 0
	for {
		//准备要发送的字符串
		// msg := fmt.Sprintf("Hello World, %03d", i)
		n, err := conn.Write(buf)
		if err != nil {
			println("Write Buffer Error:", err.Error())
			break
		}
		// fmt.Printf("write: %s\n", msg)
		if n != len(buf) {
			fmt.Printf("write error len: %v", n)
			return
		}

		//从服务器端收字符串
		n, err = conn.Read(buf)
		if err != nil {
			println("Read Buffer Error:", err.Error())
			break
		}
		// fmt.Printf("read: %s\n", string(buf[0:n]))
		if n != len(buf) {
			fmt.Printf("write error len: %v", n)
		}

		//等一秒钟
		// time.Sleep(time.Second)

		if count%10000 == 0 {
			fmt.Printf("time: %v, count: %v\n", time.Now(), count)
			count = 0
		}
		count += 1
	}
}
