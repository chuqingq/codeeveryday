package main

import "testing"

// dynamic

func BenchmarkNonGenericAdd(b *testing.B) {
	a := &AdderStruct{}
	for i := 0; i < b.N; i++ {
		doNonGeneric(a)
	}
}

type Adder interface {
	Add(a int, b int) int
}

type AdderStruct struct {
}

func (astruct *AdderStruct) Add(a int, b int) int {
	return a + b
}

func doNonGeneric(addstruct Adder) {
	res := addstruct.Add(1, 2)
	if res != 3 {
		panic("11111:")
	}
}

// generic

func BenchmarkGenericAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doGeneric()
	}
}

func addGeneric[T int | float32 | float64](a T, b T) T {
	return a + b
}

func doGeneric() {
	res := addGeneric(1, 2)
	if res != 3 {
		panic("11111:")
	}
}

/*
chuqq@chuqq-r7000p/m/d/t/p/c/g/20221007-generic-bench $ go test -bench .
goos: linux
goarch: amd64
pkg: generic_bench
cpu: AMD Ryzen 7 5800H with Radeon Graphics
BenchmarkNonGenericAdd-4        849662235                1.416 ns/op
BenchmarkGenericAdd-4           848209634                1.422 ns/op
PASS
ok      generic_bench   2.700s
*/
