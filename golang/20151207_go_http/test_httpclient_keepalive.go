package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	client := &http.Client{Transport: &http.Transport{
		MaxIdleConnsPerHost: 1,
	}}

	_, err := client.Post("http://localhost:58080/hello", "application/octet-stream", strings.NewReader("request body"))
	if err != nil {
		log.Fatalf("post1 error: %v\n", err)
	}
	log.Printf("post1 ok\n")

	time.Sleep(15 * time.Second)

	_, err = client.Post("http://localhost:58080/hello2", "application/octet-stream", strings.NewReader("request body"))
	if err != nil {
		log.Fatalf("post2 error: %v\n", err)
	}
	log.Printf("post2 ok\n")

	_, err = client.Post("http://localhost:58080/hello3", "application/octet-stream", strings.NewReader("request body"))
	if err != nil {
		log.Fatalf("post3 error: %v\n", err)
	}
	log.Printf("post3 ok\n")

	time.Sleep(15 * time.Second)
}
