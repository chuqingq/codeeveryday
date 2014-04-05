package main

import (
	// "fmt"
	// "net"
)

func main() {
	c := make(chan string)

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
}
