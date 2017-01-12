package main

import (
	"os"
	"bufio"
	"log"
	"net/http"
	"io/ioutil"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatalf("open file error: %s\n", err.Error())
	}
	defer f.Close()

	r := bufio.NewReader(f)

	req, err := http.ReadRequest(r)
	if err != nil {
		log.Fatalf("http read request error: %s\n", err.Error())
	}

	
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("http read request error: %s\n", err.Error())
	}
	log.Printf("body: %s\n", string(b))

	log.Printf("content-length: %s\n", req.Header.Get("Content-Length"))

}