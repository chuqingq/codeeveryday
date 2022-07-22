package main

import (
	"fmt"
	"time"
	"strconv"
	// "strings"
)

func main() {
	t1 := time.Now().UnixMilli()
	arr := make(map[string]int64, 1000000)
	// var b strings.Builder
	for i := 0; i < 1000000; i++ {
		// b.Reset()
		value := time.Now().Unix()
		// key := strconv.FormatInt(int64(i), 10) + "_" + strconv.FormatInt(value, 10)
		// key := fmt.Sprintf("%v_%v", i, value)
		// key := strconv.Itoa(int(i)) + "_" + strconv.Itoa(int(value))
		// fmt.Fprintf(&b, "%v_%v", i, value)
		k := strconv.Itoa(int(i)) + "_" + strconv.Itoa(int(value))
		arr[k] = value
	}
	t2 := time.Now().UnixMilli()
	fmt.Printf("%d ms, count: %v\n", t2 - t1, len(arr))
}

// go build -o hash3go hash3.go
// 287 ms, count: 1000000
