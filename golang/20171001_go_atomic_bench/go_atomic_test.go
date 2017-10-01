package main

import (
	"testing"
	"sync/atomic"
)

func BenchmarkAtomicAdd(b *testing.B) {
	var value int32
	for i :=0; i < b.N; i++ {
		atomic.AddInt32(&value, 1)
	}
}

func BenchmarkIntAdd(b *testing.B) {
	var value int32
	for i :=0; i < b.N; i++ {
                value++
        }
}

