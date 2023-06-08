package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	defer log.Printf("end")

	// for {
	// 	time.Sleep(time.Second)
	// }

	sigChan := make(chan os.Signal, 8)
	signal.Notify(sigChan)
	sig := <-sigChan
	log.Printf("接受到来自系统的信号：%v", sig)
}
