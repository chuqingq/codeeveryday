package main

func hello() {
	println("hello")
}

func main() {
	go hello()
}

// 不会打印出内容
