package main

import (
	"log"
	"net"
	"time"
)

const COUNT = 100000

func main() {
	// 创建连接
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		// 172.26.20.58
		IP:   net.IPv4(127,   0,   0,  1),
		Port: 8080,
	})
	if err != nil {
		log.Println("DialUDP error: %v", err)
		return
	}
	defer socket.Close()

	// 发送数据
	senddata := []byte("PING!")
	data := make([]byte, 1024)

	start := time.Now()
	for i := 0; i < COUNT; i++ {
		writen, err := socket.Write(senddata)
		if err != nil {
			log.Printf("Write error: %v", err)
			return
		}
		if writen != len(senddata) {
			log.Printf("writen[%v] != len(senddata)[%v]", writen, len(senddata))
			return
		}
	
		// 接收数据
		readn, _, err := socket.ReadFromUDP(data)
		if err != nil {
			log.Printf("ReadFromUDP: %v", err)
			return
		}

		if readn != writen {
			log.Printf("readn[%v] != writen[%v]", readn, writen)
			return
		}
		// fmt.Printf("%v: %s\n", remoteAddr, data[0:read])
	}
	log.Printf("%v", time.Now().Sub(start)/COUNT)
}

// local output:
// 2018/02/09 09:13:58 23.287µs
// remote output:
// 2018/02/09 09:20:45 353.44µs
// udp pingpong：在wsl中，45us；在vb ubuntu-mdc中，329us。
