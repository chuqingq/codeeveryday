package main

import "fmt"

//#include "mytest.h"
import "C"

//export MyAddGo
func MyAddGo(a, b int) int {
    return a+b
}

func main() {
    res := C.MyAddC(1, 2)
    fmt.Printf("res: %d\n", res)
}
