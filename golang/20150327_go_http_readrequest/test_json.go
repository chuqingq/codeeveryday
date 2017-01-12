package main

import (
    "fmt"
    "encoding/json"
)

type MyStruct struct {
    Test1 int `json:"test1"`
    Test2 int `json:"test2"`
}

func main() {
    data := []byte(`
    {"test1":1, "test2":2}
    `)
    
    result := make(map[string]interface{})
    err := json.Unmarshal(data, &result)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Printf("%+v\n", result)
}
