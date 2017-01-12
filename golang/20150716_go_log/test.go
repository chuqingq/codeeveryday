package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// 如果直接用os.Create()，不会自动sync
	file, err := os.Create("test.log")
	// file, err := os.OpenFile("test.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("create file error: %v\n", err)
		return
	}

	logger := log.New(file, "", log.LstdFlags|log.Lshortfile)
	logger.Printf("hello world\n")
	// file.Sync()

	err = http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("listenAndServe: ", err)
	}
}
