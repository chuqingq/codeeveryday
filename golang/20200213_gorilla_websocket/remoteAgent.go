package main

import (
	"log"
	"strings"

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

	// // 接收mainControl的消息
	// msg, err := conn.Recv()
	// if err != nil {
	// 	log.Printf("receive error: %v", err)
	// 	return
	// }
	// log.Printf("recv msg: %v", msg)

	// 等待接收启动请求
	msg, err := conn.Recv()
	log.Printf("recv: %v, err: %v", msg, err)

	// 回复启动的端口。这里直接写IP
	msg.Cmd = "start res"
	msg.IP = "123"
	conn.Send(msg)
	log.Printf("send %v", msg)

	// 等待接收下一条消息
	msg, err = conn.Recv()
	log.Printf("recv: %v, err: %v", msg, err)
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
	/*
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
	*/

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
	
	log.Printf("send register")

	msg, err = c.Recv()
	log.Printf("recv %v, err: %v", msg, err)

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
	c.wsConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	return c.wsConn.Close()
}
