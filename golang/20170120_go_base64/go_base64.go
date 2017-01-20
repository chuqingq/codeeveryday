package main

import (
	"encoding/base64"
	"log"
)

func main() {
	src := []byte{1, 2, 3, 4}
	log.Printf("src: %v", src)

	encoded := base64.StdEncoding.EncodeToString(src)
	log.Printf("encoded: %v", encoded)

	src1, err := base64.StdEncoding.DecodeString(encoded)
	log.Printf("src1: %v, err: %v", src1, err)
}

