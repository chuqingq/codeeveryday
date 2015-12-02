package main

import (
    "net/http"
    "fmt"
    "ioutil"
    "log"
    "encoding/json"
)

type mystruct struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func main() {
    res, err := http.Get("http://www.baidu.com")
    if err != nil {
        println("error: ", err)
        return
    }
    defer res.Body.Close()

    content, err := ioutil.ReadAll(res.Body)
    if err != nil {
    	println("error")
    	return
    }

    var str mystruct
    _ = json.UnMarshal(content, &str)

    fmt.Printf("%#v\n", res)
}