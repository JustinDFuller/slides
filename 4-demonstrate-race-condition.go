package main

import "fmt"

type WaitGroup struct {
	i int
}

func (wg *WaitGroup) Add(i int) {
	wg.i += i
}

func (wg *WaitGroup) Done() {
	wg.i -= 1
}

func (wg *WaitGroup) Wait() {
	for {
		fmt.Println("WaitGroup Count: ", wg.i)
		if wg.i <= 0 {
			break
		}
	}
}

func main() {
	var wg WaitGroup

	for i := 0; i < 1000; i++ { // HL
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("Saw i =", i)
		}(i)
	}

	wg.Wait()
	fmt.Println("Done")
}
