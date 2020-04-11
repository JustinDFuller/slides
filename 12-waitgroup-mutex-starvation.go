package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

func main() {
	var wg WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("WaitGroup.Wait() got the lock", i, "times.")
}

var i int64

func (wg *WaitGroup) Wait() {
	for {
		wg.lock.Lock()
		atomic.AddInt64(&i, 1)
		if wg.i <= 0 {
			wg.lock.Unlock()
			break
		}
		wg.lock.Unlock()
	}
}

type WaitGroup struct {
	i    int
	lock Mutex // HL
}

func (wg *WaitGroup) Add(i int) {
	defer wg.lock.Unlock() // HL
	wg.lock.Lock()         // HL
	wg.i += i
}

func (wg *WaitGroup) Done() {
	defer wg.lock.Unlock() // HL
	wg.lock.Lock()         // HL
	wg.i -= 1
}

const (
	unlocked = iota
	locked
)

type Mutex struct {
	isLocked int32
}

func (m *Mutex) Unlock() {
	atomic.CompareAndSwapInt32(&m.isLocked, locked, unlocked) // HL
	runtime.Gosched()
}

func (m *Mutex) Lock() {
	for {
		if atomic.CompareAndSwapInt32(&m.isLocked, unlocked, locked) { // HL
			// The mutex was unlocked, now it's locked.
			return
		}
	}
}
