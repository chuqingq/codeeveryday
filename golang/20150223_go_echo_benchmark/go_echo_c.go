package main

import (
	"fmt"
	"net"
	"time"
)

const COUNT = 100000

func main() {
	buf := make([]byte, 10, 10)
	PING := []byte("PING!")

	before := time.Now()
	c, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic("Dial error: " + err.Error())
	}

	for i := 0; i < COUNT; i++ {
		n, e := c.Write(PING)
		if e != nil {
			panic("write error: " + e.Error())
		}
		if n != 5 {
			panic("write len not 5")
		}

		n, e = c.Read(buf)
		if e != nil {
			panic("read error: " + e.Error())
		}
		if 5 != n {
			panic("read len not 5")
		}
	}

	fmt.Printf("%v\n", time.Now().Sub(before))
}
