package main

import (
	"fmt"
	"sync"
)

type Counter interface {
	Increment(workerID int)
	Value() int
}

type SafeCounter struct {
	mu      sync.Mutex
	counter int
}

func (sc *SafeCounter) Increment(workerID int) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.counter++
	fmt.Printf("Worker %d increased the counter to %d\n", workerID, sc.counter)
}

func (sc *SafeCounter) Value() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.counter
}
