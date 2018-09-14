package main

import (
    "net/http"
    "fmt"
    "log"
    "time"
)

var response = []byte("hello world")

func HandleFuncToken(res http.ResponseWriter, req * http.Request) {
    res.Write([]byte("this is token"))
}

func HandleFunc(res http.ResponseWriter, req * http.Request) {
    log.Printf("receive")
    if pusher, ok := res.(http.Pusher); ok {
        //支持Push
        time.Sleep(3*time.Second)
        if err := pusher.Push("/token", nil); err != nil {
            log.Printf("Failed to push: %v", err)
        }
    } else {
        log.Printf("not pusher")
    }
    res.Write(response)
}

func main() {
    http.HandleFunc("/", HandleFunc)
    http.HandleFunc("/upstream/token", HandleFuncToken)
    err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
    if err != nil {
        fmt.Println(err)
    }
}
