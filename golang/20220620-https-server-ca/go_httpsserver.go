package main

import (
    "net/http"
    "fmt"
    "log"
)

var response = []byte("hello world")


func HandleFunc(res http.ResponseWriter, req * http.Request) {
    log.Printf("receive")
    res.Write(response)
}

func main() {
    http.HandleFunc("/", HandleFunc)
    err := http.ListenAndServeTLS(":443", "http.crt", "http.key", nil)
    if err != nil {
        fmt.Println(err)
    }
}
