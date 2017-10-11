package main

import (
	"testing"
)

func BenchmarkHello0(b *testing.B) {
	Hello()
}
/*
$ go test -bench Hello -benchmem -memprofile mem.out -benchtime 0s
goos: linux
goarch: amd64
BenchmarkHello-8   	       1	     37969 ns/op	  106496 B/op	       1 allocs/op
PASS
ok  	_/home/chuqq/temp/codeeveryday/golang/20170928_go_benchmark	0.002s
*/

func BenchmarkHello1(b *testing.B) {
	Hello2()
}

