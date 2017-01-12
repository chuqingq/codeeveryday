package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.360kb.com/kb/2_150.html")
	if err != nil {
		log.Fatalf("http.Get error: %s\n", err.Error())
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ioutil.ReadAll error: %s\n", err.Error())
	}

	log.Printf("recv: %s\n", string(buf))
}
