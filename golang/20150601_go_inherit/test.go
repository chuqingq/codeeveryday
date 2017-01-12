package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) SetName(n string) {
	p.name = n
}

type Student struct {
	Person

	name string
	id   int
}

func (s *Student) Name() string {
	return "student " + s.Person.Name() + " " + s.name
}

// func (s *Student) SetName() string
func main() {
	p := &Student{}
	p.name = "name22"
	p.SetName("name1")
	p.age = 20
	p.id = 100
	fmt.Printf("name=%s\n", p.Name())
	fmt.Printf("%v", p)
}
