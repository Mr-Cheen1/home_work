package main

import (
	"encoding/json"
	"fmt"
)

// Структура Book.
type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float64 `json:"rate"`
}

// Функция реализует интерфейс Marshaller для структуры Book.
func (b *Book) MarshalJSON() ([]byte, error) {
	type Alias Book

	if b.Title == "" {
		return nil, fmt.Errorf("book title cannot be empty")
	}

	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	})
}

// UnmarshalJSON реализует интерфейс Unmarshaller для структуры Book.
func (b *Book) UnmarshalJSON(data []byte) error {
	type Alias Book
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	}

	// Десериализация данных во временную структуру.
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Проверка, что название книги не пустое после десериализации.
	if b.Title == "" {
		return fmt.Errorf("book title cannot be empty after deserialization")
	}

	return nil
}

// Функция выполняет сериализацию слайса объектов Book в JSON.
func MarshalJSONSlice(books []Book) ([]byte, error) {
	for _, book := range books {
		if book.Title == "" || book.Author == "" {
			return nil, fmt.Errorf("the book must have a title and author")
		}
	}
	return json.Marshal(books)
}

// Функция выполняет десериализацию слайса объектов Book из JSON.
func UnmarshalJSONSlice(data []byte) ([]Book, error) {
	var books []Book
	if err := json.Unmarshal(data, &books); err != nil {
		return nil, err
	}

	// Проверка каждой книги на наличие названия и автора после десериализации.
	for _, book := range books {
		if book.Title == "" || book.Author == "" {
			return nil, fmt.Errorf("found a book without title or author after deserialization")
		}
	}

	return books, nil
}

func main() {
	// Определение слайса структур Book
	books := []Book{
		{
			ID:     1,
			Title:  "Test Book 1",
			Author: "Author 1",
			Year:   2023,
			Size:   123,
			Rate:   4.5,
		},
		{
			ID:     2,
			Title:  "Test Book 2",
			Author: "Author 2",
			Year:   2024,
			Size:   456,
			Rate:   4.7,
		},
	}

	// Сериализация в JSON
	booksJSON, err := MarshalJSONSlice(books)
	if err != nil {
		fmt.Println("Slice marshaling error:", err)
		return
	}
	fmt.Println("Books in JSON:", string(booksJSON))

	// Десериализация из JSON
	var newBooks []Book
	newBooks, err = UnmarshalJSONSlice(booksJSON)
	if err != nil {
		fmt.Println("Slice unmarshaling error:", err)
		return
	}
	fmt.Printf("Books from JSON: %+v\n", newBooks)
}
