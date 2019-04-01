package main

import "C"

//export Add
func Add(a int, b int) int {
    return a + b
}

func main() {
}

