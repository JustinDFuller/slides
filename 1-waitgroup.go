package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup // HL

	for i := 0; i < 10; i++ {
		wg.Add(1) // HL
		go func(i int) {
			defer wg.Done() // HL
			fmt.Println("Saw i =", i)
		}(i)
	}

	wg.Wait() // HL
	fmt.Println("Done")
}
