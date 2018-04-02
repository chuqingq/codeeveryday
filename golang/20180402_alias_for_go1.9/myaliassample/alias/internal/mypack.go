package internal

type MyPack struct {
	a int
	b int
}

func (mypack *MyPack) SetA(a int) {
	mypack.a = a
}

func (mypack *MyPack) SetB(b int) {
	mypack.b = b
}

func (mypack *MyPack) Add() int {
	return mypack.a + mypack.b
}
