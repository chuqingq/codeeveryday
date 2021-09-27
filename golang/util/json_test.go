package util

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	sjson "github.com/bitly/go-simplejson"
)

type A struct {
	data interface{}
}

func TestJsonUnmarshal(t *testing.T) {
	b := []byte(`{"A":1}`)
	// m := &Message{}
	m := &sjson.Json{}
	json.Unmarshal(b, m)
	fmt.Printf("json.Json: %v", m)
	m2 := &A{}
	json.Unmarshal(b, m2)
	fmt.Printf("A: %v", m2)
}

// TODO TestJsonGet
// TODO TestJsonSet
// TODO TestJsonMap

func TestJsonArray(t *testing.T) {
	m := NewJson()
	d := []interface{}{1, 2, 3}
	m.SetPath("", d)
	s, _ := m.EncodePretty()
	log.Printf("m: %v", m)
	log.Printf("json: %v", string(s))
	// log.Printf("%v", m.GetPath(""))
	// 测试Array方法
	ms := m.MustArray()
	// ms[2].SetPath([]string{}, 100)
	for _, mm := range ms {
		log.Printf("%v", mm.MustInt(100))
	}
}
