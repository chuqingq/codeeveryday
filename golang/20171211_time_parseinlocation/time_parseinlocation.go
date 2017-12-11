package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    str := t.Format("2006-01-02 15:04:05")
    location, _ := time.LoadLocation("Asia/Shanghai")
    t2, _ := time.ParseInLocation("2006-01-02 15:04:05", str, location)
    fmt.Printf("%v\n", t2) // 2017-12-11 09:29:07 +0800 CST
}

