
package main

import (
	"log"
	// "net"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	// "time"
)

var (
	n   = 100000
	c   = 1
	url = "http://127.0.0.1:8081"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < c; i++ {
		wg.Add(1)
		go func() {
			worker()
			wg.Done()
		}()
	}

	wg.Wait()
}

func worker() {
	client := &http.Client{}
	for i := 0; i < n/c; i++ {
		// res, err := client.Get(url)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Printf("newrequest error: %v", err)
			continue
		}
		// req.Header.Add("Connection", "Keep-Alive")
		res, err := client.Do(req)
		if err != nil {
			log.Printf("err: %v", err)
			continue
		}
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
	}
	log.Printf("concurrent finish")
}

