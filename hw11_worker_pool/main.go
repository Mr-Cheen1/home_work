package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func worker(workerID int, counter Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Increment(workerID)
}

func main() {
	workers := 5
	counter := &SafeCounter{}

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(i, counter, &wg)
	}

	wg.Wait()
	fmt.Println("All workers have completed their work. Total value of the counter:", counter.Value())
}
