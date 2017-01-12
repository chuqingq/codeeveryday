package main

import (
	. "fmt"
	"github.com/go-vgo/robotgo"
)

func main() {
	x, y := robotgo.GetMousePos()
	Println("pos:", x, y)
}

