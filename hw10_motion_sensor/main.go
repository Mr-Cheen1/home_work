package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	rawDataChan := readSensorData()
	processedDataChan := make(chan float64)

	go processDataUsingAverage(rawDataChan, processedDataChan)

	for data := range processedDataChan {
		fmt.Printf("Processed data: %.2f\n", data)
	}
}

// Функция имитации считывания данных с сенсора и передачи их в канал.
func readSensorData() <-chan int {
	dataChan := make(chan int)
	go func() {
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
	}()
	return dataChan
}

// Функция для обработки данных с использованием среднего значения.
func processDataUsingAverage(rawDataChan <-chan int, processedDataChan chan<- float64) {
	defer close(processedDataChan)

	var sum int
	var count int

	for data := range rawDataChan {
		sum += data
		count++
		if count == 10 {
			average := float64(sum) / 10
			processedDataChan <- average
			sum = 0
			count = 0
		}
	}
}
