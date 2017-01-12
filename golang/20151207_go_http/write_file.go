package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("logic.aof")
	if err != nil {
		log.Fatalf("open file error: %v\n", err)
	}

	f.WriteString("$10\264q\000\0003132333435363711\000\000\000\000$10\264q\000\0003132333435363738\000\000\000\000$10\264q\000\0003132333435363714\000\000\000\000")
}
