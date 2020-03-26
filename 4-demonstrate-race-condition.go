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
		if wg.i <= 0 {
			break
		}
	}
}

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
