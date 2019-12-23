package main

import (
	"math/rand"
	"sync"
)

func main() {
	var total int
	var wg sync.WaitGroup

	for i := 0; i<10; i++ {
		wg.Add(1)
		go func() {
			for j:= 0; j < 1000; j++ {
				total += readNumber()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func readNumber() int {
	return rand.Intn(10)
}