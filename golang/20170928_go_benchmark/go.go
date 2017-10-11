package main

func Hello() {
	_ = make([]byte, 10240, 10240)
	_ = make([]byte, 102400, 102400)
}

type myStruct1 struct {
	bytes [10240]byte
}

type myStruct2 struct {
	bytes [10240000]byte
}

func Hello1() {
	_ = &myStruct1{}
	_ = &myStruct2{}
}

func Hello2() {
	_ = new(myStruct2)
}

