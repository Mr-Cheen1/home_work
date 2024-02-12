package main

import (
	"sync"
	"testing"
)

type MockCounter struct {
	Counter
	IncrementCalls int
}

func (mc *MockCounter) Increment(_ int) {
	mc.IncrementCalls++
}

func TestWorker(t *testing.T) {
	mockCounter := &MockCounter{}
	var localWg sync.WaitGroup
	localWg.Add(1)

	worker(1, mockCounter, &localWg)

	localWg.Wait()

	if mockCounter.IncrementCalls != 1 {
		t.Errorf("Expected Increment to be called once, called %d times", mockCounter.IncrementCalls)
	}
}
