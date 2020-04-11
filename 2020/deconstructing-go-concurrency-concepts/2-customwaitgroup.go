package main

import "fmt"

func main() {
	var wg WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("Saw i =", i)
		}(i)
	}

	wg.Wait()
	fmt.Println("Done")
}

type WaitGroup struct{}

func (wg *WaitGroup) Add(i int) {

}

func (wg *WaitGroup) Done() {

}

func (wg *WaitGroup) Wait() {

}
