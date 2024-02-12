package main

import (
	"sync"
	"testing"
)

func TestCounterValue(t *testing.T) {
	counter := &SafeCounter{}
	var localWg sync.WaitGroup
	workers := 3
	localWg.Add(workers)

	for i := 0; i < workers; i++ {
		go worker(i, counter, &localWg)
	}

	localWg.Wait()

	expectedValue := workers
	if counter.Value() != expectedValue {
		t.Errorf("Expected counter value to be %d, got %d", expectedValue, counter.Value())
	}
}

func TestCounterThreadSafety(t *testing.T) {
	counter := &SafeCounter{}
	var localWg sync.WaitGroup
	workers := 1000
	localWg.Add(workers)

	for i := 0; i < workers; i++ {
		go worker(i, counter, &localWg)
	}

	localWg.Wait()

	if counter.Value() != workers {
		t.Errorf("Expected counter value to be %d after %d workers, got %d", workers, workers, counter.Value())
	}
}

func TestIntegrationOfWorkersAndSafeCounter(t *testing.T) {
	counter := &SafeCounter{}
	var localWg sync.WaitGroup
	workers := 10
	localWg.Add(workers)

	for i := 0; i < workers; i++ {
		go worker(i, counter, &localWg)
	}

	localWg.Wait()

	expectedValue := workers
	if counter.Value() != expectedValue {
		t.Errorf("Expected counter value to be %d after %d workers, got %d", expectedValue, workers, counter.Value())
	}
}
