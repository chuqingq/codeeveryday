package main

import (
    "net/http"
    "fmt"
)

func main() {
    res, err := http.Get("http://www.baidu.com")
    if err != nil {
        println("error: ", err)
        return
    }

    fmt.Println(res)
}