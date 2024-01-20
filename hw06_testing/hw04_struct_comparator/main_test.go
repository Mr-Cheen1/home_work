package main

import (
	"testing"
)

func TestCompareBooks(t *testing.T) {
	tests := []struct {
		name     string
		mode     Comparator
		book1    Book
		book2    Book
		expected bool
	}{
		{
			name:     "Сравнение по году, книга1 > книга2",
			mode:     Year,
			book1:    Book{year: 2000},
			book2:    Book{year: 1999},
			expected: true,
		},
		{
			name:     "Сравнение по размеру, книга1 < книга2",
			mode:     Size,
			book1:    Book{size: 100},
			book2:    Book{size: 200},
			expected: false,
		},
		{
			name:     "Сравнение по рейтингу, книга1 == книга2",
			mode:     Rate,
			book1:    Book{rate: 4.5},
			book2:    Book{rate: 4.5},
			expected: false,
		},
		{
			name:     "Сравнение по году, книга1 < книга2",
			mode:     Year,
			book1:    Book{year: 1998},
			book2:    Book{year: 1999},
			expected: false,
		},
		{
			name:     "Неверный режим сравнения",
			mode:     Comparator(-1),
			book1:    Book{year: 2000},
			book2:    Book{year: 1999},
			expected: false,
		},
		{
			name:     "Сравнение по размеру, книга1 > книга2",
			mode:     Size,
			book1:    Book{size: 300},
			book2:    Book{size: 200},
			expected: true,
		},
		{
			name:     "Сравнение по рейтингу, книга1 > книга2",
			mode:     Rate,
			book1:    Book{rate: 5.0},
			book2:    Book{rate: 4.5},
			expected: true,
		},
		{
			name:     "Сравнение по рейтингу, книга1 < книга2",
			mode:     Rate,
			book1:    Book{rate: 3.5},
			book2:    Book{rate: 4.0},
			expected: false,
		},
		{
			name:     "Сравнение по году с одинаковыми годами",
			mode:     Year,
			book1:    Book{year: 2000},
			book2:    Book{year: 2000},
			expected: false,
		},
		{
			name:     "Сравнение по размеру с одинаковыми размерами",
			mode:     Size,
			book1:    Book{size: 200},
			book2:    Book{size: 200},
			expected: false,
		},
		{
			name:     "Сравнение с неверными данными в книге1",
			mode:     Year,
			book1:    Book{year: -1},
			book2:    Book{year: 2000},
			expected: false,
		},
		{
			name:     "Сравнение с неверными данными в книге2",
			mode:     Size,
			book1:    Book{size: 100},
			book2:    Book{size: -50},
			expected: true,
		},
		{
			name:     "Сравнение по рейтингу с неверными данными в обеих книгах",
			mode:     Rate,
			book1:    Book{rate: -1.0},
			book2:    Book{rate: -2.0},
			expected: false,
		},
		{
			name:     "Сравнение по году с неверными данными в обеих книгах",
			mode:     Year,
			book1:    Book{year: -100},
			book2:    Book{year: -200},
			expected: true,
		},
		{
			name:     "Неопределенный компаратор, книга1 > книга2",
			mode:     Comparator(100),
			book1:    Book{size: 300},
			book2:    Book{size: 200},
			expected: false,
		},
		{
			name:     "Сравнение книг с нулевыми значениями",
			mode:     Year,
			book1:    Book{year: 0},
			book2:    Book{year: 0},
			expected: false,
		},
		{
			name:     "Сравнение книг с максимальными значениями года",
			mode:     Year,
			book1:    Book{year: 9999},
			book2:    Book{year: 9999},
			expected: false,
		},
		{
			name:     "Сравнение книг с одинаковыми значениями во всех полях",
			mode:     Size,
			book1:    Book{size: 100, rate: 5.0, year: 2000},
			book2:    Book{size: 100, rate: 5.0, year: 2000},
			expected: false,
		},
		{
			name:     "Сравнение книги с неинициализированными значениями и книги с валидными значениями",
			mode:     Rate,
			book1:    Book{},
			book2:    Book{rate: 4.5},
			expected: false,
		},
		{
			name:     "Сравнение книг, где книга1 имеет отрицательный размер",
			mode:     Size,
			book1:    Book{size: -100},
			book2:    Book{size: 200},
			expected: false,
		},
		{
			name:     "Сравнение книг, где книга1 имеет отрицательный рейтинг, а книга2 - положительный",
			mode:     Rate,
			book1:    Book{rate: -5.0},
			book2:    Book{rate: 4.5},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CompareBooks(tt.mode, tt.book1, tt.book2)
			if result != tt.expected {
				t.Errorf("CompareBooks(%v, %v, %v) = %v, expected value %v", tt.mode, tt.book1, tt.book2, result, tt.expected)
			}
		})
	}
}
