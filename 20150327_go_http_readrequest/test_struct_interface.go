package main

import "fmt"

type MyInterface interface {
    Add(int)
}

type MyStruct struct {
    A int
    MyInterface
}

func (s *MyStruct)Add(a int) {
    s.A += a
}

func main() {
    m := MyStruct{}
    m.A = 1
    m.Add(2)
    fmt.Printf("%d\n", m.A)
}