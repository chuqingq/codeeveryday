package main

import (
	"log"

	json "github.com/bitly/go-simplejson"
)

// 验证go-simplejson支持按级别设置

func main() {
	j := json.New()
	j.SetPath([]string{"a", "b", "c"}, "c-data")
	k := json.New()
	k.SetPath([]string{"a", "b"}, "b-data")

	j.SetPath([]string{"a", "b", "d"}, k)
	s, _ := j.Encode()
	log.Printf("j: %v", string(s))

	b := j.GetPath("a", "b")
	s, _ = b.Encode()
	log.Printf("b: %v", string(s))
}
