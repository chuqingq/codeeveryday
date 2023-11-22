package main

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const CmdSetNetwork string = "set_network" // 设置网络

func TestMain(t *testing.T) {
	// 设备启动
	dev := &Device{
		SerialNumber: "1234567890",
		IP:           "192.168.0.140",
		Port:         12345,
	}

	err := dev.Start()
	if err != nil {
		log.Printf("Register error: %v", err)
		return
	}
	defer dev.Stop()

	// 设备接收消息
	go func() {
		for cmd := range dev.RecvChan {
			if cmd.Cmd == Req(CmdSetNetwork) {
				// do setnetwork
				log.Printf("response setnetwork")
				cmd.Cmd = Res(CmdSetNetwork)
				dev.Send(*cmd)
			}
		}
	}()

	time.Sleep(100 * time.Millisecond)

	// 搜索工具启动
	cli := &Client{}

	err = cli.Start()
	if err != nil {
		log.Printf("Discover error: %v", err)
		return
	}
	defer cli.Stop()

	// 发送discover请求
	cmd := &Command{
		Cmd: Req(CmdDiscover),
	}
	cli.Send(*cmd)

	// client接收新消息
	const newIP = "192.168.0.141"
	for cmd := range cli.RecvChan {
		if cmd.Cmd == Res(CmdDiscover) {
			log.Printf("client recv discover res: %v", cmd)
			cmd.Cmd = Req(CmdSetNetwork)
			cmd.Device.IP = newIP
			cmd.Device.Port = 12346
			cli.Send(*cmd)
		} else if cmd.Cmd == Res(CmdSetNetwork) {
			log.Printf("client recv setnetwork res: %v", cmd)
			assert.Equal(t, newIP, cmd.Device.IP)
			return
		}
	}
}
