package main

import "fmt"

// 只需添加这句import即可

import "C"

func main() {
	fmt.Println("vim-go")
}

// 缩小体积 go build -ldflags "-s" go_dynamically_linked.go
