package main

import (
	"log"
)

func main() {
	min := uint64(1 - 1)
	max := uint64(18446744073709551615)

	var a uint64 = uint64(min) - 1

	var b uint64 = uint64(max) + 1

	log.Printf("a=%d, b=%d\n", a, b)
}

// 2015/07/02 09:05:37 a=18446744073709551615, b=0
