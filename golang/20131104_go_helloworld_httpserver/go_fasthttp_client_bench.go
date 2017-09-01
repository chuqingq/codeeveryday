
package main

import (
	"log"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

var (
	n   = 1000000
	c   = 100
	url = "http://127.0.0.1:8081"
)

func main() {
	var wg sync.WaitGroup
	start := time.Now()

	count := n/c
	for i := 0; i < c; i++ {
		wg.Add(1)
		go func() {
			worker(count)
			wg.Done()
		}()
	}
	log.Printf("started...")

	wg.Wait()
	elapsed := time.Now().Sub(start).Seconds()
	log.Printf("speed: %v requests/second", float64(n)/elapsed)
}

func worker(count int) {
	client := &fasthttp.Client{}
	buf := make([]byte, 128)
	for i := 0; i < count; i++ {
		_, _, err := client.Get(buf, url)
		if err != nil {
			log.Printf("err: %v", err)
		}
	}
}

