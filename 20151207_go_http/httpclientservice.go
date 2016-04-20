package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net"
	"net/http"
)

func main() {
	port := flag.Int("port", 56544, "listen udp port")
	flag.Parse()

	if *port < 0 || *port > 65535 {
		log.Fatalf("listen udp port[%d] invalid\n", *port)
	}

	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: *port,
	})
	if err != nil {
		log.Fatalf("listen udp port[%d] error: %v\n", err)
	}
	defer socket.Close()

	data := make([]byte, 4096)
	for {
		n, err := socket.Read(data)
		if err != nil {
			log.Printf("udp read error: %v\n", err)
			continue
		}

		var req Request
		err = json.Unmarshal(data[:n], &req)
		if err != nil {
			log.Printf("json decode error: %v\n", err)
			continue
		}

		go sendRequest(&req)
	}
}

type Request struct {
	Method string      `json:"method"`
	Url    string      `json:"url"`
	Data   interface{} `json:"data"`
}

func sendRequest(req *Request) {
	data, err := json.Marshal(req.Data)
	if err != nil {
		log.Printf("json encode data error: %v\n", err)
		return
	}

	if req.Method == "" {
		if len(data) == 0 {
			req.Method = "GET"
		} else {
			req.Method = "POST"
		}
	}

	client := &http.Client{}
	request, _ := http.NewRequest(req.Method, req.Url, bytes.NewReader(data))
	response, err := client.Do(request)
	if err != nil {
		log.Printf("http request[%s %s] error: %v\n", req.Method, req.Url, err)
		return
	}
	defer response.Body.Close()

	log.Printf("http request[%s %s] status: %v\n", response.Status)
}
