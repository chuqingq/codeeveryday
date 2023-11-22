package main

import (
	"log"
	"net"
	"sync"
)

type MulticastUDP struct {
	Conn     *net.UDPConn
	RecvChan chan []byte
}

func (s *MulticastUDP) Start() error {
	var err error
	// 创建UDP连接
	s.Conn, err = net.ListenMulticastUDP("udp", nil, GetMulticastAddr())
	if err != nil {
		log.Printf("ListenMulticastUDP error: %v", err)
		return err
	}

	// 接收数据
	s.RecvChan = make(chan []byte, 8)
	go func() {
		defer close(s.RecvChan)
		buf := make([]byte, 1024*1024)
		for {
			n, _, err := s.Conn.ReadFromUDP(buf)
			if err != nil {
				log.Printf("ReadFromUDP error: %v", err)
				return
			}
			// log.Printf("Received: %s\n", buf[:n])
			// 拷贝并发送
			s.RecvChan <- append([]byte(nil), buf[:n]...)
		}
	}()

	return nil
}

func (s *MulticastUDP) Stop() {
	s.Conn.Close()
	// close(s.RecvChan)
}

func (s *MulticastUDP) Send(data []byte) error {
	// 连接
	conn, err := net.DialUDP("udp", nil, GetMulticastAddr())
	if err != nil {
		return err
	}

	// 发送数据
	_, err = conn.Write(data)
	if err != nil {
		log.Printf("Send error: %v", err)
		return err
	} else {
		log.Printf("socket writeto success: %v", string(data))
	}
	return nil
}

var MulticastAddr string = "224.0.0.1:9999"
var multicastAddr *net.UDPAddr
var onceMulticastAddr sync.Once

func GetMulticastAddr() *net.UDPAddr {
	onceMulticastAddr.Do(func() {
		var err error
		multicastAddr, err = net.ResolveUDPAddr("udp", MulticastAddr)
		if err != nil {
			log.Printf("ResolveUDPAddr [%v] error: %v", MulticastAddr, err)
		}
	})
	return multicastAddr
}
