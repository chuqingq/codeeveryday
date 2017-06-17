// 如果method的receiver是针对*T的，那通过T调用这个method，效果一样
package main

import (
	"fmt"
)

type Data struct {
	x int
}

func (self Data) ValueTest(a int) { // func ValueTest(self Data);
	self.x = a
	fmt.Printf("Value: %p\n", &self)
}
func (self *Data) PointerTest(a int) { // func PointerTest(self *Data);
	self.x = a
	fmt.Printf("Pointer: %p\n", self)
}
func main() {
	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)
	d.ValueTest(1) // ValueTest(d)
	fmt.Printf("%+v\n", d)
	d.PointerTest(2) // PointerTest(&d)
	fmt.Printf("%+v\n", d)
	p.ValueTest(3) // ValueTest(*p)
	fmt.Printf("%+v\n", p)
	p.PointerTest(4) // PointerTest(p)
	fmt.Printf("%+v\n", p)
}
