package main

import (
	"log"
	"net"
	"os"
	"strings"

	"golang.org/x/net/ipv4"
)

var multicastAddr, _ = net.ResolveUDPAddr("udp", "224.0.0.1:9999")

func main() {
	if len(os.Args) > 1 && strings.Contains(os.Args[1], "server") {
		server()
		select {}
	} else {
		client()
	}
}

func client() {
	// 单播发送
	laddr, _ := net.ResolveUDPAddr("udp", ":0")
	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		log.Fatalf("DialUDP error: %v", err)
	}

	_, err = conn.WriteTo([]byte("hello"), multicastAddr)
	if err != nil {
		log.Printf("WriteTo error: %v", err)
	}

	var buf [1024]byte
	n, raddr, err := conn.ReadFromUDP(buf[:])
	if err != nil {
		log.Printf("ReadFromUDP error: %v", err)
	} else {
		log.Printf("recv %v from %v", string(buf[:n]), raddr)
	}
}

func server() {
	// multicast
	// // 方案1：windows10本地收不到
	// multicast, err := net.ListenMulticastUDP("udp", nil, multicastAddr)
	// if err != nil {
	// 	log.Fatalf("ListenMulticastUDP error: %v", err)
	// }

	// 方案2：支持win10本地收发
	multicast, err := net.ListenUDP("udp4", multicastAddr)
	if err != nil {
		log.Fatalf("ListenUDP error: %v", err)
	}
	pc := ipv4.NewPacketConn(multicast)
	var inf *net.Interface
	if err := pc.JoinGroup(inf, multicastAddr); err != nil {
		log.Printf("joinGroup error: %v. interface: %v", err, inf)
	}
	if loop, err := pc.MulticastLoopback(); err == nil {
		log.Printf("MulticastLoopback status:%v", loop)
		if !loop {
			if err := pc.SetMulticastLoopback(true); err != nil {
				log.Printf("SetMulticastLoopback error:%v", err)
				// continue
			}
		}
	}

	// 接收
	go func() {
		var buf [1024]byte
		for {
			n, raddr, err := multicast.ReadFromUDP(buf[:])
			if err != nil {
				log.Fatalf("multicast.ReadFromUDP error: %v", err)
			}
			log.Printf("multicast recv: %v from %v", string(buf[:n]), raddr)

			_, err = multicast.WriteTo([]byte("world"), raddr)
			if err != nil {
				log.Printf("respond error: %v", err)
			}
		}
	}()
}
