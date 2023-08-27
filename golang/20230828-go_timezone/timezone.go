package main

import (
	"time"
)

func main() {
	GetVersion()
}

func GetVersion() string {
	// Sun 27 Aug 2023 02:13:45 PM UTC
	nowStr := "2023-08-27T14:13:45Z"
	now, _ := time.Parse(time.RFC3339, nowStr)

	loc, _ := time.LoadLocation("Asia/Shanghai")
	nowStr2 := now.In(loc).Format(time.RFC3339)
	println(nowStr2)
	return nowStr2
}
