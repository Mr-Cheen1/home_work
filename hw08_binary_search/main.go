package main

import (
	"errors"
	"fmt"
)

// Функция BinarySearch реализует алгоритм двоичного поиска,
// которая возвращает ошибку, если срез пуст,
// меньше или больше первого и последнего значения в массиве.
func BinarySearch(slice []int, target int) (int, error) {
	if len(slice) == 0 {
		return -1, errors.New("slice is empty")
	}

	if target < slice[0] || target > slice[len(slice)-1] {
		return -1, errors.New("target is out of bounds")
	}

	left, right := 0, len(slice)-1
	for left <= right {
		mid := left + (right-left)/2
		midValue := slice[mid]
		switch {
		case midValue == target:
			return mid, nil
		case midValue < target:
			left = mid + 1
		default:
			right = mid - 1
		}
	}

	if left > right {
		return -1, errors.New("target not found")
	}

	return -1, nil
}

func main() {
	slice := []int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	target := 11

	index, err := BinarySearch(slice, target)
	switch {
	case err != nil:
		fmt.Println("Error:", err)
	case index != -1:
		fmt.Printf("Element found at index: %d\n", index)
	default:
		fmt.Println("Element not found in the slice.")
	}
}
