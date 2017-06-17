package main

import (
	"log"
)

type mystr struct {
	a int
}

func (m *mystr) Add(b int) {
	m.a += b
}

func myfunc(callback func(b int)) {
	callback(10)
}

func main() {
	a := &mystr{}
	myfunc(a.Add)
	log.Printf("%v", a.a)
}

//   2017/06/16 15:37:07 10
