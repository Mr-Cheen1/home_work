package main

import (
	"testing"
	"time"
)

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
	dataChan := readSensorData()

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

func TestProcessDataUsingAverageMaxValues(t *testing.T) {
	rawDataChan := make(chan int, 10)
	processedDataChan := make(chan float64, 1)

	go func() {
		for i := 0; i < 10; i++ {
			rawDataChan <- 100
		}
		close(rawDataChan)
	}()

	go processDataUsingAverage(rawDataChan, processedDataChan)

	select {
	case result := <-processedDataChan:
		expected := 100.0
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	case <-time.After(2 * time.Second):
		t.Error("Test timed out")
	}
}

func TestProcessDataUsingAverageWithFewDataPoints(t *testing.T) {
	rawDataChan := make(chan int, 5)
	processedDataChan := make(chan float64)

	go func() {
		for i := 1; i <= 5; i++ {
			rawDataChan <- i
		}
		close(rawDataChan)
	}()

	go processDataUsingAverage(rawDataChan, processedDataChan)

	timeout := time.After(2 * time.Second)
	select {
	case _, ok := <-processedDataChan:
		if ok {
			t.Error("Did not expect to receive processed data due to insufficient input data points")
		}
	case <-timeout:
	}
}

func TestReadSensorDataGeneratesData(t *testing.T) {
	dataChan := readSensorData()
	timeout := time.After(2 * time.Second)
	select {
	case _, ok := <-dataChan:
		if !ok {
			t.Error("Expected data channel to be open and receive data, but it was closed")
		}
	case <-timeout:
		t.Error("Expected to receive data, but did not within the timeout period")
	}
}
