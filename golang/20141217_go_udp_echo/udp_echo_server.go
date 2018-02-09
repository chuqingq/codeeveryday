package main

import (
	"log"
	"net"
)

func main() {
	// 创建监听
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	})
	if err != nil {
		log.Println("ListenUDP error: %v", err)
		return
	}
	defer socket.Close()

	data := make([]byte, 4096)
	for {
		// 读取数据
		readn, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			log.Printf("ReadFromUDP error: %v", err)
			continue
		}
		// fmt.Printf("%v: %s\n\n", remoteAddr, data[0:read])

		// 发送数据
		// senddata := []byte("hello client!")
		writen, err := socket.WriteToUDP(data[:readn], remoteAddr)
		if err != nil {
			log.Printf("WriteToUDP error: %v", err)
			return
		}

		if writen != readn {
			log.Printf("writen[%v] != readn[%v]", writen, readn)
			return
		}
	}
}
