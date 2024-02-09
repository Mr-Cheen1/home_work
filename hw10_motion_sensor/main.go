package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	rawDataChan := make(chan int)
	processedDataChan := make(chan float64)
	go readSensorData(rawDataChan)
	go processDataUsingAverage(rawDataChan, processedDataChan)

	for data := range processedDataChan {
		fmt.Printf("Processed data: %.2f\n", data)
	}
}

// Функция имитации считывания данных с сенсора и передачи их в канал.
func readSensorData(dataChan chan<- int) {
	defer close(dataChan)
	for start := time.Now(); time.Since(start) < time.Minute; {
		n, err := rand.Int(rand.Reader, big.NewInt(100))
		if err != nil {
			fmt.Println("Random number generation error:", err)
			return
		}
		dataChan <- int(n.Int64())
		time.Sleep(time.Second)
	}
}

// Функция для обработки данных с использованием среднего значения.
func processDataUsingAverage(rawDataChan <-chan int, processedDataChan chan<- float64) {
	defer close(processedDataChan)
	dataBatch := make([]int, 0, 10)

	for data := range rawDataChan {
		dataBatch = append(dataBatch, data)
		if len(dataBatch) == 10 {
			average := calculateAverage(dataBatch)
			processedDataChan <- average
			dataBatch = dataBatch[:0]
		}
	}
}

// Функция для вычисления среднего значения.
func calculateAverage(data []int) float64 {
	if len(data) == 0 {
		return 0.0
	}
	var sum int
	for _, value := range data {
		sum += value
	}
	return float64(sum) / float64(len(data))
}
