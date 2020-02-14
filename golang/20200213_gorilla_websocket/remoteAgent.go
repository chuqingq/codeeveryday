package main

import (
	"log"
	"net"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	// 连接mainControl
	url := "ws://127.0.0.1:8080/agent"
	mykey := "chuqingqing13"
	conn, err := ConnectAndRegisterToMainControl(url, mykey)
	if err != nil {
		log.Printf("connect mainControl error: %v", err)
		return
	}

	// 关闭
	defer conn.Close()

	// 接收mainControl的消息
	msg, err := conn.Recv()
	if err != nil {
		log.Printf("receive error: %v", err)
		return
	}
	log.Printf("recv msg: %v", msg)

	msg, err = conn.Recv()
	if err != nil {
		log.Printf("recv error: %v", err)
		return
	}
	log.Printf("recv msg: %v", msg)
}

type MainControlConn struct {
	url    string
	wsConn *websocket.Conn
}

type MainControlMsg struct {
	Cmd string
	Key string
	IP  string
}

func ConnectAndRegisterToMainControl(url string, mykey string) (*MainControlConn, error) {
	wsConn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	c := &MainControlConn{
		url:    url,
		wsConn: wsConn,
	}

	// TCP keepalive
	netConn := wsConn.UnderlyingConn()
	tcpConn, ok := netConn.(*net.TCPConn)
	if !ok {
		log.Printf("keepalive error: %v", err)
	}
	err = tcpConn.SetKeepAlive(true)
	if err != nil {
		wsConn.Close()
		return nil, err
	}
	err = tcpConn.SetKeepAlivePeriod(2 * time.Minute)
	if err != nil {
		wsConn.Close()
		return nil, err
	}

	// 注册
	localAddr := wsConn.LocalAddr().String()
	ip := strings.Split(localAddr, ":")[0]
	log.Printf("local ip: %v", ip)
	msg := &MainControlMsg{
		Cmd: "register",
		Key: mykey,
		IP:  ip,
	}
	err = c.Send(msg)
	if err != nil {
		wsConn.Close()
		return nil, err
	}

	return c, nil
}

func (c *MainControlConn) GetWebSocket() *websocket.Conn {
	return c.wsConn
}

func (c *MainControlConn) Send(msg *MainControlMsg) error {
	return c.wsConn.WriteJSON(msg)
}

func (c *MainControlConn) Recv() (*MainControlMsg, error) {
	msg := &MainControlMsg{}
	err := c.wsConn.ReadJSON(msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (c *MainControlConn) Close() error {
	// TODO?
	c.wsConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	return c.wsConn.Close()
}
