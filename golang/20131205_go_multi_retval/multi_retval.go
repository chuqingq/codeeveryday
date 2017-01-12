package main

import "fmt"

func init() {
	fmt.Println("init")
}

func fa() (int, int) {
	return 1,2
}

func fb() (int, int) {
	return 3,4
}

func main() {
	a,b := fa()
	fmt.Printf("a=%d, b=%d\n", a, b)
	c,b := fb()
	fmt.Printf("c=%d, b=%d\n", c, b)
}
