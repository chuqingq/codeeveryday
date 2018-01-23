package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.22.41:8888")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	buf := []byte("hello world")
	count := 0
	start := time.Now().UnixNano()
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
			println("Write len error:", n)
		}

		//从服务器端收字符串
		n, err = conn.Read(buf[:])
		if err != nil {
			println("Read Buffer Error:", err.Error())
			break
		}

		if n != len(buf) {
			println("Read len error:", n)
		}

		// fmt.Printf("read: %s\n", string(buf[0:n]))
		// fmt.Printf("%v", time.Now().UnixNano()-start)
		count++
		if count % 10000 == 0 {
			now := time.Now().UnixNano()
			fmt.Printf("elapsed: %v, avg: %v\n", now-start, (now-start)/10000)
			start = now
			count = 0
		}
	}
}

// 单次
// linux-qiwr:~ # ./tcp_echo_client 
// read: hello world
// 505230

// 循环
// linux-qiwr:~ # ./tcp_echo_client 
// elapsed: 2452126940, avg: 245212
// elapsed: 2452734333, avg: 245273
// elapsed: 2658226987, avg: 265822
// elapsed: 2656539814, avg: 265653
// elapsed: 2451967858, avg: 245196
// elapsed: 2656135622, avg: 265613
// elapsed: 2454530446, avg: 245453
// elapsed: 2452213075, avg: 245221
// elapsed: 2653260511, avg: 265326
// elapsed: 2658530878, avg: 265853

// linux-qiwr:~ # ping 192.168.22.40
// PING 192.168.22.40 (192.168.22.40) 56(84) bytes of data.
// 64 bytes from 192.168.22.40: icmp_seq=1 ttl=64 time=0.026 ms
// 64 bytes from 192.168.22.40: icmp_seq=2 ttl=64 time=0.027 ms
// 64 bytes from 192.168.22.40: icmp_seq=3 ttl=64 time=0.026 ms
// 64 bytes from 192.168.22.40: icmp_seq=4 ttl=64 time=0.027 ms
// 64 bytes from 192.168.22.40: icmp_seq=5 ttl=64 time=0.022 ms

