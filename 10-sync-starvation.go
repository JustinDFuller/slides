package main

import (
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

func (m *Mutex) Unlock() {
	atomic.CompareAndSwapInt32(&m.isLocked, locked, unlocked) // HL
}

func (m *Mutex) Lock() {
	for {
		if atomic.CompareAndSwapInt32(&m.isLocked, unlocked, locked) { // HL
			// The mutex was unlocked, now it's locked.
			return
		}
	}
}

func main() {
	done := make(chan bool, 1)
	var mu Mutex // HL

	go func() {
		for {
			select {
			case <-done:
				return
			default:
				mu.Lock() // HL
				time.Sleep(100 * time.Microsecond)
				mu.Unlock() // HL
			}
		}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Microsecond)
		mu.Lock()   // HL
		mu.Unlock() // HL
	}
	done <- true
}
