package main

import (
	"log"
	"net"
	//"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func main() {
	conn, err := net.ListenPacket("ip4:udp", "0.0.0.0")
	if err != nil {
		log.Fatalf("ListenPacket error: %v\n", err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Printf("ReadFrom error: %v\n", err)
		}

		// if addr.String() != "192.168.23.72" {
		// 	continue
		// }

		if addr.String() != "127.0.0.1" {
			continue
			log.Printf("ReadFrom %v\n", addr)
		}

		reqPacket := gopacket.NewPacket(buf[:n], layers.LayerTypeUDP, gopacket.Default)
		udpLayer := reqPacket.Layer(layers.LayerTypeUDP)
		if udpLayer == nil {
			log.Printf("udpLayer nil")
			continue
		}
		req := udpLayer.(*layers.UDP)
		
		// 计算目的端端口
		dstPort := uint16(buf[2])*256+uint16(buf[3])
		log.Printf("ReadFrom %v, dstport: %v, %v\n", addr, dstPort, string(buf[8:n]))
		if dstPort != 12345 {
			continue
		}
		/*
		// 交换源端和目的端端口
		buf[0], buf[1], buf[2], buf[3] = buf[2], buf[3], buf[0], buf[1]
		conn.WriteTo(buf[:n], addr)
		*/
	}
}

/*
$ sudo ./rawsock_udp 
2017/08/11 16:09:24 ReadFrom 127.0.0.1, 24, 123123123123123

$ nc -u 127.0.0.1 12345
123123123123123

*/
