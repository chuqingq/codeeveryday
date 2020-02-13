package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	// 启动HTTP服务
	addr := "127.0.0.1:8080"
	// 添加agent的websocket处理
	http.HandleFunc("/agent", agentHandler)
	log.Fatal(http.ListenAndServe(addr, nil))
}

var upgrader = websocket.Upgrader{}

func agentHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	var msg MainControlMsg
	for {
		err := c.ReadJSON(&msg)
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", msg)

		// 恢复响应
		res := &MainControlMsg{
			Cmd: "register ok",
		}
		err = c.WriteJSON(res)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

type MainControlMsg struct {
	Cmd string
	Key string
	IP  string
}
