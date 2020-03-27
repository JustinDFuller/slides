package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var wg WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("Saw i =", i)

			if i%10 == 0 {
				defer wg.Done() // HL
				wg.Add(1)       // HL
				fmt.Println("Added an extra")
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Done")
}

func (wg *WaitGroup) Wait() {
	for {
		fmt.Println("WaitGroup Count: ", wg.i)
		wg.lock.Lock()
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
	atomic.CompareAndSwapInt32(&m.isLocked, locked, unlocked)
}

func (m *Mutex) Lock() {
	for {
		if atomic.CompareAndSwapInt32(&m.isLocked, unlocked, locked) {
			// The mutex was unlocked, now it's locked.
			return
		}
	}
}
