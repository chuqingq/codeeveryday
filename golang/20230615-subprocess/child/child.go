package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/chuqingq/go-util"
)

func main() {
	// 开启stderr
	go func() {
		for {
			log.Printf("this is child log: %v", time.Now())
			time.Sleep(time.Second)
		}
	}()

	// 从stdin读数据，写入stdout
	decoder := json.NewDecoder(os.Stdin)
	encoder := json.NewEncoder(os.Stdout)
	for {
		m := util.NewMessage()
		err := decoder.Decode(m)
		if err != nil {
			log.Printf("read error: %v", err)
			return
		}
		encoder.Encode(m)
	}
}
