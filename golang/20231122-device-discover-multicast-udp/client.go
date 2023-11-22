package main

import (
	"encoding/json"
	"log"
)

// Client 设备发现工具客户端。下发发现请求，接收发现回复。
type Client struct {
	Socket   *MulticastUDP
	RecvChan chan *Command
}

// Start 客户端启动发现。返回recvChan
func (s *Client) Start() error {
	s.Socket = &MulticastUDP{}
	err := s.Socket.Start()
	if err != nil {
		return err
	}

	s.RecvChan = make(chan *Command, 2)
	// 接收socket的消息，并转成Command发给RecvChan
	go func() {
		log.Printf("client is recving...")
		defer log.Printf("client socket recv finish")
		defer close(s.RecvChan)
		for b := range s.Socket.RecvChan {
			cmd := &Command{}
			err := json.Unmarshal(b, &cmd)
			if err != nil {
				log.Printf("client json unmarshal error: %v", err)
				continue
			}
			s.RecvChan <- cmd
		}
	}()
	return nil
}

// Send 客户端发送/回复命令
func (s *Client) Send(cmd Command) error {
	b, err := json.Marshal(cmd)
	if err != nil {
		return err
	}
	return s.Socket.Send(b)
}

// Stop 客户端注销
func (s *Client) Stop() {
	s.Socket.Stop()
}
