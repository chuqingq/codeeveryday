package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	t1 := time.Now().UnixNano()
	arr := make(map[string]int64)
	for i := 0; i < 1000000; i++ {
		value := time.Now().Unix()
		// key := strconv.FormatInt(int64(i), 10) + "_" + strconv.FormatInt(value, 10)
		// key := fmt.Sprintf("%v_%v", i, value)
		key := strconv.Itoa(int(i)) + "_" + strconv.Itoa(int(value))
		arr[key] = value
	}
	t2 := time.Now().UnixNano()
	fmt.Printf("%d ms\n", (t2 - t1)/1000000)
}

