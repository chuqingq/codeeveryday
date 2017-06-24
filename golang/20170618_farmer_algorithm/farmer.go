// 最初有5个农民采金子，每个农民每秒采8块钱，每个农民50块钱。问：怎么最快采到1000块钱
package main

import (
	"fmt"
)

// 如果目标是farmer个农民，第几秒能采到1000块钱
func caijin(farmer, monkey int) int {
	curFarmer := 5
	curMonkey := 0
	for i := 1; true; i++ {
		// 这秒的钱
		curMonkey += curFarmer * 8
		// 如果钱超过1000，则返回成功
		if curMonkey >= monkey {
			return i
		}
		// 如果有余钱，且不到目标农民数，则尽早买农民
		for curMonkey >= 50 && curFarmer < farmer {
			curMonkey -= 50
			curFarmer += 1
		}
		fmt.Printf("i: %v, curMonkey: %v, curFarmer: %v\n", i, curMonkey, curFarmer)
	}
	return -1
}

func main() {
	// min := 125
	// for i := 5; i < 10000; i++ {
	// 	if caijin(i, 1000) < min {
	// 		min = caijin(i, 1000)
	// 		fmt.Printf("==== %v: %v\n", i, min)
	// 	}
	// }
	// 12:17
	fmt.Printf("==== %v: %v\n", 12, caijin(12, 1000))
}
