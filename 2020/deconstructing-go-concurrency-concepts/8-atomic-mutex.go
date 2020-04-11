package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var wg WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("Saw i =", i)
		}(i)
	}

	wg.Wait()
	fmt.Println("Done")
}

func (wg *WaitGroup) Wait() {
	for {
		wg.lock.Lock()
		fmt.Println("WaitGroup Count: ", wg.i)
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
}

func (m *Mutex) Lock() {
	for {
		if atomic.CompareAndSwapInt32(&m.isLocked, unlocked, locked) { // HL
			// The mutex was unlocked, now it's locked.
			return
		}
	}
}
