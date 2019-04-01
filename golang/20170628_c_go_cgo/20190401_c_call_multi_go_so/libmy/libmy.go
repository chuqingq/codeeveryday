package main

import "net/http"
import "C"

//export Hello
func Hello() string {
        http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("hello world"))
        })
        http.ListenAndServe(":8080", nil)

    return "Hello"
}

func main() {
}
