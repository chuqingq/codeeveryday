package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	for {
		a, b := test()
		fmt.Printf("a:%v, b:%v\n", a, b)
		time.Sleep(time.Second)
	}
}

func test() (r int, err error) {
	r, err = 0, errors.New("something wrong")

	defer func() {
		fmt.Println("defer")
		if err := recover(); err != nil {
			fmt.Printf("recover: %v\n", err)
		}
	}()

	a := 0
	return 1 / a, nil
}

// 每秒输出
// defer
// recover: runtime error: integer divide by zero
// a:0, b:something wrong
