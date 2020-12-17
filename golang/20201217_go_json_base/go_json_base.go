package main

import (
	"encoding/json"
	"fmt"
)

type Base struct {
	Base int
}

func (b *Base) ToJson() {
	fmt.Printf("base: %p\n", b)
	p := interface{}(b)
	b1, _ := json.Marshal(p)
	fmt.Printf("base.json: %v\n", string(b1))
}

func (b *Base) ToJson2(p interface{}) {
	b1, _ := json.Marshal(p)
	fmt.Printf("base.json2: %v\n", string(b1))
}

type Child struct {
	Base
	Child int
}

func main() {
	var c Child
	fmt.Printf("child: %p\n", &c)
	c.ToJson()
	c.ToJson2(&c)
}
