package main

import (
	"log"
	"net/http"
	// "time"

	"github.com/gorilla/websocket"
)

func main() {
	// 启动HTTP服务
	addr := "127.0.0.1:8080"
	// 添加agent的websocket处理
	http.HandleFunc("/metisweb", metisWebHandler)
	http.HandleFunc("/agent", agentHandler)
	log.Printf("listen and serve at %v", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

var upgrader = websocket.Upgrader{}

type myConn struct {
	wsConn *websocket.Conn
	ch chan *MainControlMsg
}

var hostsInfo = make(map[string]*myConn)

func agentHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("connected")

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	var msg MainControlMsg

	// 接收register请求
	err = c.ReadJSON(&msg)
	log.Printf("recv: %v, err: %v", msg, err)


	// 回复register响应
	res := &MainControlMsg{
		Cmd: "register ok",
	}
	err = c.WriteJSON(res)
	log.Printf("send register ok: %v", err)

	
	ch := make(chan *MainControlMsg)
	// 保存映射关系
	hostsInfo[msg.Key] = &myConn {
		wsConn: c,
		ch: ch,
	}

	// // 注册关闭回调：从map中删除 // 不需要，同步ReadJSON()可以感知到断连
	// c.SetCloseHandler(func(code int, text string) error {
	// 	log.Printf("closed: %v", msg.Key)
	// 	delete(hostsInfo, msg.Key)
	// 	return nil
	// })

	// !!! 如果这里直接返回，c就关掉了。猜测是http做的。
	for {
		msg = MainControlMsg{}
		err = c.ReadJSON(&msg)
		log.Printf("main loop recv: %v, err: %v", msg, err)
		if err != nil {
			log.Printf("read error: %v", err)
			c.Close()
			delete(hostsInfo, msg.Key)
		}
		ch <- &msg
	}
}

type MainControlMsg struct {
	Cmd string
	Key string
	IP  string
}

func metisWebHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.FormValue("key")

	// 给客户端发个请求，获取到响应，然后再回复metis-web
	c := hostsInfo[key].wsConn
	msg := &MainControlMsg {
		Cmd: "start",
	}
	err := c.WriteJSON(msg)
	log.Printf("send %v, err: %v", msg, err)

	msg = <-hostsInfo[key].ch
	log.Printf("recv: %v", msg)

	w.Write([]byte(msg.IP))
}
