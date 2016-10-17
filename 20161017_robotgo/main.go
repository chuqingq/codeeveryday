package main

import (
	"github.com/go-vgo/robotgo"
)

func main() {
	robotgo.ScrollMouse(100, "up")
	select {}
}

