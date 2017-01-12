package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(RandInt64(100, 101))
	fmt.Println(RandInt64(100, 200))
}

func RandInt64(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min)
}
