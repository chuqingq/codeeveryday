package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	buf := make([]byte, 5, 5)

	s, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("Listen error: " + err.Error())
	}

	for {
		c, err := s.Accept()
		if err != nil {
			println("Accept error: " + err.Error())
			continue
		}
		before := time.Now()

		for {
			n, err := c.Read(buf)
			if err != nil {
				println("read error: " + err.Error())
				break
			}
			if n != 5 {
				println("read len not 5")
				break
			}
			c.Write(buf[:n])
		}

		fmt.Printf("%v\n", time.Now().Sub(before))
	}
}
