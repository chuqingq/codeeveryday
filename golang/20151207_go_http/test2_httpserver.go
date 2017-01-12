package main

import (
    "net/http"
    "fmt"
)

func HandleFunc(res http.ResponseWriter, req * http.Request) {
    fmt.Fprint(res, "hello world")
}

func main() {
    http.HandleFunc("/hello", HandleFunc)
    err := http.ListenAndServe(":8081", nil)
    if err != nil {
        fmt.Println(err)
    }
}
