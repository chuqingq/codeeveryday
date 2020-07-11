package main

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

// https://colobu.com/2017/03/09/implement-TryLock-in-Go/

const mutexLocked = 1 << iota

type Mutex struct {
	sync.Mutex
}

func (m *Mutex) TryLock() bool {
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked)
}
