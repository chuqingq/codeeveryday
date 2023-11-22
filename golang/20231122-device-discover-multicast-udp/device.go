package main

import (
	"encoding/json"
	"log"
)

// Start 设备注册
func (s *Device) Start() error {
	s.Socket = &MulticastUDP{}
	err := s.Socket.Start()
	if err != nil {
		return err
	}

	s.RecvChan = make(chan *Command, 2)
	// 从socket接收消息，并处理
	go func() {
		log.Printf("device is recving...")
		defer log.Printf("Device socket recv finish")
		defer close(s.RecvChan)

		for b := range s.Socket.RecvChan {
			// log.Printf("Device recv %v", string(b))
			cmd := &Command{}
			err := json.Unmarshal(b, &cmd)
			if err != nil {
				log.Printf("Device recv json parse error: %v", err)
				continue
			}
			// 处理cmd
			if cmd.Cmd == Req(CmdDiscover) {
				// 发现设备
				log.Printf("Device recv discover req: %v", cmd)
				cmd.Cmd = Res(CmdDiscover)
				cmd.Device = *s
				s.Send(*cmd)
			} else {
				s.RecvChan <- cmd
			}
		}
	}()

	return nil
}

// Send 设备发送/回复命令
func (s *Device) Send(cmd Command) error {
	b, err := json.Marshal(cmd)
	if err != nil {
		return err
	}
	return s.Socket.Send(b)
}

// Stop 设备注销
func (s *Device) Stop() {
	s.Socket.Stop()
}
