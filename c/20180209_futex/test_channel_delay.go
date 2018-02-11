package main

import (
	"log"
	"time"
)

func main() {
	log.Printf("starting...")
	ch := make(chan bool, 1)
	go func() {
		<-ch
		log.Printf("%v", time.Now().UnixNano())
	}()

	time.Sleep(2 * time.Second)
	log.Printf("%v", time.Now().UnixNano())
	ch <- true
	time.Sleep(2 * time.Second)
}

// macos:
// $ go run test_channel_delay.go
// 2018/02/11 19:19:19 starting...
// 2018/02/11 19:19:21 1518347961360782000
// 2018/02/11 19:19:21 1518347961360941000
