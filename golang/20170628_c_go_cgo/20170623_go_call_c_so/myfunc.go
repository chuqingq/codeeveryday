package main

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./ -lmyfunc
#include "myfunc.h"
*/
import "C"

// 上面的cgo的注释，必须紧挨着import "C"，中间不能有空行

func main() {
	C.hello()
}
