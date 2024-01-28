package main

import (
	"testing"
)

func TestBinarySearchEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		target   int
		expected int
		wantErr  bool
	}{
		{
			name:     "Reverse ordered slice",
			slice:    []int{17, 15, 13, 11, 9, 7, 5, 3, 1},
			target:   11,
			expected: -1,
			wantErr:  true,
		},
		{
			name:     "Large slice",
			slice:    makeLargeSlice(),
			target:   123456,
			expected: 123456,
			wantErr:  false,
		},
		{
			name:     "Target at the start",
			slice:    []int{1, 2, 3, 4, 5},
			target:   1,
			expected: 0,
			wantErr:  false,
		},
		{
			name:     "Target at the end",
			slice:    []int{1, 2, 3, 4, 5},
			target:   5,
			expected: 4,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinarySearch(tt.slice, tt.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s: BinarySearch() error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("%s: BinarySearch() = %v, want %v", tt.name, got, tt.expected)
			}
		})
	}
}

// makeLargeSlice создает большой срез для тестирования производительности.
func makeLargeSlice() []int {
	largeSlice := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		largeSlice[i] = i
	}
	return largeSlice
}

func TestBinarySearchSpecialCases(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		target   int
		expected int
		wantErr  bool
	}{
		{
			name:     "Element in the middle",
			slice:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			target:   5,
			expected: 4,
			wantErr:  false,
		},
		{
			name:     "Even number of elements",
			slice:    []int{1, 2, 3, 4, 5, 6},
			target:   3,
			expected: 2,
			wantErr:  false,
		},
		{
			name:     "Odd number of elements",
			slice:    []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
			wantErr:  false,
		},
		{
			name:     "Element less than all in slice",
			slice:    []int{10, 20, 30, 40, 50},
			target:   5,
			expected: -1,
			wantErr:  true,
		},
		{
			name:     "All elements greater than zero",
			slice:    []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
			wantErr:  false,
		},
		{
			name:     "All elements less than zero",
			slice:    []int{-5, -4, -3, -2, -1},
			target:   -3,
			expected: 2,
			wantErr:  false,
		},
		{
			name:     "Slice with nil or zero values",
			slice:    []int{0, 0, 0, 0, 0},
			target:   0,
			expected: 2,
			wantErr:  false,
		},
		{
			name:     "Max int32 value",
			slice:    []int{0, 1, 2, 3, 2147483647},
			target:   2147483647,
			expected: 4,
			wantErr:  false,
		},
		{
			name:     "Min int32 value",
			slice:    []int{-2147483648, -1, 0, 1, 2},
			target:   -2147483648,
			expected: 0,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinarySearch(tt.slice, tt.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s: BinarySearch() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
			if got != tt.expected {
				t.Errorf("%s: BinarySearch() = %v, want %v", tt.name, got, tt.expected)
			}
		})
	}
}

func TestBinarySearchErrorHandling(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		target   int
		expected int
		wantErr  bool
	}{
		{
			name:     "Unsorted slice",
			slice:    []int{10, 7, 3, 8, 2},
			target:   3,
			expected: -1,
			wantErr:  true,
		},
		{
			name:     "Repeating elements with target present",
			slice:    []int{2, 3, 3, 3, 4},
			target:   3,
			expected: 2,
			wantErr:  false,
		},
		{
			name:     "Empty slice",
			slice:    []int{},
			target:   3,
			expected: -1,
			wantErr:  true,
		},
		{
			name:     "Single element slice, target not present",
			slice:    []int{1},
			target:   3,
			expected: -1,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinarySearch(tt.slice, tt.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s: BinarySearch() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
			if got != tt.expected {
				t.Errorf("%s: BinarySearch() = %v, want %v", tt.name, got, tt.expected)
			}
		})
	}
}
