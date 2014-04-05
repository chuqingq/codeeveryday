package main

import (
	"time"
	// "fmt"
	// "net"
)

func main() {
	c := make(chan string)
	defer close(c)

	go func() {
		// for {
			data := <- c
			println("server recv", data)
			c <- "world"
		// }
	}()
	
	c <- "hello"
	res := <- c
	println("client recv", res)

	select {
		case res := <- c:
			println("client recv", res)
		case <- time.After(time.Second * 2):
			println("timeout")
	}
}
