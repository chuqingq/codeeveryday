package mytest

import (
	"testing"
	"sync"
	"sync/atomic"
)

func TestAbc(t *testing.T) {
	m1 := map[string]string{}
	m2 := m1
	m1["1"] = "1"
	m1["2"] = "2"
	println(len(m2))
	// t.Fail()
}

func BenchmarkHello(b *testing.B) {
	var a int64
	var mutex sync.Mutex
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mutex.Lock()
		a += 1
		mutex.Unlock()
	}
	b.StopTimer()
	if a != int64(b.N) {
		b.Fail()
	}
}

func BenchmarkWorld(b *testing.B) {
	var a int64
	var c int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c = atomic.LoadInt64(&a)
		for !atomic.CompareAndSwapInt64(&a, c, c+1) {
			c = atomic.LoadInt64(&a)
		}
	}
	b.StopTimer()
	if a != int64(b.N) {
		b.Fail()
	}
}

/*
$ go test -bench=.*
2
BenchmarkHello-8   	100000000	        14.9 ns/op
BenchmarkWorld-8   	200000000	         8.84 ns/op
PASS
ok  	_/home/chuqq/temp/codeeveryday/c/20170710_memory_order/mytest	4.981s
*/
