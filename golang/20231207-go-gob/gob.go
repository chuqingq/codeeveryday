package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func main() {
	data := []byte("Hello, 世界")

	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	enc.Encode(data)
	fmt.Printf("buffer: %v\n", buf.Len())

	dec := gob.NewDecoder(&buf)

	decode := make([]byte, 1024)
	err := dec.Decode(&decode)

	fmt.Printf("decode: %v, %v, %v ", err, string(decode), len(decode))
}

/*
buffer: 17
decode: <nil>, Hello, 世界, 13
*/
