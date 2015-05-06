package main

import "fmt"


type MyStruct struct {
    A int
    B int
}

func main() {
    m := MyStruct{A: 10, B:20}
    fmt.Printf("%d\n", m.A)
}