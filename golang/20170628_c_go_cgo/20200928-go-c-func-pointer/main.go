package main

/*
 #cgo CFLAGS: -I .
 #cgo LDFLAGS: -L . -lclibrary
 #include "clibrary.h"
 int callOnMeGo_cgo(int in); // 声明
*/
import "C"

import (
	"fmt"
	"unsafe"
)

//export callOnMeGo
func callOnMeGo(in int) int {
	return in + 1
}

func main() {
	fmt.Printf("Go.main(): calling C function with callback to us\n")

	//使用unsafe.Pointer转换
	C.some_c_func((C.callback_fcn)(unsafe.Pointer(C.callOnMeGo_cgo)))
}
