package main

import (
	"log"
	
	"myaliassample/alias"
)

func main() {
    mystruct := &alias.MyStruct{}
	mystruct.SetA(1)
	mystruct.SetB(2)
	log.Printf("result: %d", mystruct.Add())
}
