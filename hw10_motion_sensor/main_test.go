package main

import (
	"testing"
	"time"
)

func TestCalculateAverage(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected float64
	}{
		{"empty slice", []int{}, 0.0},
		{"single element", []int{10}, 10.0},
		{"multiple elements", []int{1, 2, 3, 4, 5}, 3.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateAverage(tt.data); got != tt.expected {
				t.Errorf("calculateAverage() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestProcessDataUsingAverage(t *testing.T) {
	rawDataChan := make(chan int, 10)
	processedDataChan := make(chan float64, 1)

	go func() {
		rawDataChan <- 1
		rawDataChan <- 2
		rawDataChan <- 3
		rawDataChan <- 4
		rawDataChan <- 5
		rawDataChan <- 6
		rawDataChan <- 7
		rawDataChan <- 8
		rawDataChan <- 9
		rawDataChan <- 10
		close(rawDataChan)
	}()

	go processDataUsingAverage(rawDataChan, processedDataChan)

	select {
	case result := <-processedDataChan:
		expected := 5.5
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	case <-time.After(2 * time.Second):
		t.Error("Test timed out")
	}
}

func TestReadSensorData(t *testing.T) {
	dataChan := make(chan int)
	go readSensorData(dataChan)

	time.Sleep(2 * time.Second)

	select {
	case data := <-dataChan:
		if data < 0 || data > 100 {
			t.Errorf("Generated data out of expected range: %d", data)
		}
	case <-time.After(1 * time.Second):
		t.Error("No data received, expected at least one data point")
	}
}
