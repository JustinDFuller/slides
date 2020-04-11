package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("Saw i =", i)
		}(i)
	}

	fmt.Println("Done")
}
