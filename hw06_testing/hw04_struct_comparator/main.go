package main

import (
	"fmt"
)

type Book struct {
	isbn   string
	title  string
	author string
	year   int
	size   int
	rate   float64
}

func (b *Book) SetISBN(isbn string) {
	b.isbn = isbn
}

func (b *Book) ISBN() string {
	return b.isbn
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) SetRate(rate float64) {
	b.rate = rate
}

func (b *Book) Rate() float64 {
	return b.rate
}

type Comparator int

const (
	Year Comparator = iota
	Size
	Rate
)

func CompareBooks(mode Comparator, b1, b2 Book) bool {
	switch mode {
	case Year:
		return b1.year > b2.year
	case Size:
		return b1.size > b2.size
	case Rate:
		if b1.rate < 0 || b2.rate < 0 {
			return false
		}
		return b1.rate > b2.rate
	default:
		return false
	}
}

func main() {
	books := []Book{
		{isbn: "1", title: "AAA", author: "A", year: 1999, size: 329, rate: 6.5},
		{isbn: "2", title: "BBB", author: "B", year: 2023, size: 213, rate: 4.0},
	}

	fmt.Println("Сравнение книг по различным критериям (год, размер, рейтинг), " +
		"вывод true если первый аргумент больше второго и false если наоборот")

	for i, book := range books {
		fmt.Printf("Книга %d: №%s, название книги: %s, автор: %s, год: %d, размер: %d, рейтинг: %.2f\n",
			i+1, book.isbn, book.title, book.author, book.year, book.size, book.rate)
	}

	fmt.Printf("Год 'Книги 1' больше года 'Книги 2': %t\n", CompareBooks(Year, books[0], books[1]))
	fmt.Printf("Размер 'Книги 1' больше размера 'Книги 2' : %t\n", CompareBooks(Size, books[0], books[1]))
	fmt.Printf("Рейтинг 'Книги 1' больше рейтинга 'Книги 2': %t\n", CompareBooks(Rate, books[0], books[1]))
}
