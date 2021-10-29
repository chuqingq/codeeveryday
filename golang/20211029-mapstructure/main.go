package main

import (
	"log"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
}

func main() {
	input := map[string]interface{}{"Name": "123123123123"}
	var p Person
	mapstructure.Decode(input, &p)
	log.Printf("%#v", p)

	m := map[string]interface{}{}
	mapstructure.Decode(p, &m)
	log.Printf("%#v", m)
}
