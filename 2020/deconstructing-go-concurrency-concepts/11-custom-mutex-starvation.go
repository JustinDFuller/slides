package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

const (
	unlocked = iota
	locked
)

type Mutex struct {
	isLocked int32
}

func (m *Mutex) Lock() {
	for {
		if atomic.CompareAndSwapInt32(&m.isLocked, unlocked, locked) { // HL
			// The mutex was unlocked, now it's locked.
			return
		}
	}
}

var i int64
var x int64

func main() {
	done := make(chan bool, 1)
	var mu Mutex

	go func() {
		for {
			select {
			case <-done:
				return
			default:
				mu.Lock()
				i++
				time.Sleep(100 * time.Microsecond)
				mu.Unlock()
			}
		}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Microsecond)
		mu.Lock()
		x++
		mu.Unlock()
	}
	done <- true
	fmt.Println("Slow got", i, "Fast got", x)
}

func (m *Mutex) Unlock() {
	atomic.CompareAndSwapInt32(&m.isLocked, locked, unlocked)
	runtime.Gosched() // HL
}
