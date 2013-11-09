package main

import (
    "net/http"
    "fmt"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir(".")))
    err := http.ListenAndServe(":8081", nil)
    if err != nil {
        fmt.Println(err)
    }
}