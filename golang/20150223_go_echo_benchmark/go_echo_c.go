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

	c, err := net.Dial("tcp", "192.168.54.118:8080")
	if err != nil {
		panic("Dial error: " + err.Error())
	}

	before := time.Now()
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

	fmt.Printf("%v\n", time.Now().Sub(before)/COUNT)
}

// local output:
// 24.688µs
// remote output:
// 385.01µs
