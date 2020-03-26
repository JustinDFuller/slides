package main

import (
	"fmt"
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

type Mutex struct {
	locked bool
}

func (m *Mutex) Unlock() {
	m.locked = false
}

func (m *Mutex) Lock() {
	if m.locked {
		for {
			if m.locked == false {
				break
			}
		}
	}

	m.locked = true
}
