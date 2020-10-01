package main

import (
	"log"

	json "github.com/bitly/go-simplejson"
)

// 验证go-simplejson支持：某个path设置为Json，按照更深层次GetPath

func main() {
	j := json.New()
	j.SetPath([]string{"a", "b", "c"}, "c-data")
	k := json.New()
	k.SetPath([]string{"a", "b"}, "b-data")

	j.SetPath([]string{"a", "b", "d"}, k.MustMap(nil) /*这里用k不行*/)
	s, _ := j.EncodePretty()
	log.Printf("j: %v", string(s))

	b := j.GetPath("a", "b", "d", "a")
	s, _ = b.EncodePretty()
	log.Printf("b: %v", string(s))

	/*
		b2 := j.GetPath("a", "b", "d")
		s, _ = b2.EncodePretty()
		log.Printf("b2: %v", string(s))
		b3 := b2.Get("a")
		s, _ = b3.EncodePretty()
		log.Printf("b3: %v", string(s))
	*/
}
