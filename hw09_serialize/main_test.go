package main

import (
	"encoding/hex"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/Mr-Cheen1/home_work/hw09_serialize/bookpb"
	"google.golang.org/protobuf/proto"
)

func TestMarshalJSONSlice(t *testing.T) {
	books := []Book{
		{
			ID:     1,
			Title:  "Test Book 1",
			Author: "Author 1",
			Year:   2023,
			Size:   100,
			Rate:   4.5,
		},
		{
			ID:     2,
			Title:  "Test Book 2",
			Author: "Author 2",
			Year:   2024,
			Size:   200,
			Rate:   4.7,
		},
	}

	booksJSON, err := MarshalJSONSlice(books)
	if err != nil {
		t.Fatalf("Marshaling error: %v", err)
	}

	var actualBooks []Book
	err = json.Unmarshal(booksJSON, &actualBooks)
	if err != nil {
		t.Fatalf("Failed to unmarshal books JSON: %v", err)
	}

	if !reflect.DeepEqual(books, actualBooks) {
		t.Errorf("Expected %v, received %v", books, actualBooks)
	}
}

func TestUnmarshalJSONSlice(t *testing.T) {
	input := `[
		{
			"id": 1,
			"title": "Test Book 1",
			"author": "Author 1",
			"year": 2023,
			"size": 100,
			"rate": 4.5
		},
		{
			"id": 2,
			"title": "Test Book 2",
			"author": "Author 2",
			"year": 2024,
			"size": 200,
			"rate": 4.7
		}
	]`
	expected := []Book{
		{
			ID:     1,
			Title:  "Test Book 1",
			Author: "Author 1",
			Year:   2023,
			Size:   100,
			Rate:   4.5,
		},
		{
			ID:     2,
			Title:  "Test Book 2",
			Author: "Author 2",
			Year:   2024,
			Size:   200,
			Rate:   4.7,
		},
	}
	books, err := UnmarshalJSONSlice([]byte(input))
	if err != nil {
		t.Errorf("Unmarshaling error: %v", err)
	}
	if !reflect.DeepEqual(books, expected) {
		t.Errorf("Expected %v, received %v", expected, books)
	}
}

// Интеграционный тест.
func TestIntegrationMarshalAndUnmarshalJSONSlice(t *testing.T) {
	originalBooks := []Book{
		{
			ID:     1,
			Title:  "Test Book 1",
			Author: "Author 1",
			Year:   2023,
			Size:   100,
			Rate:   4.5,
		},
		{
			ID:     2,
			Title:  "Test Book 2",
			Author: "Author 2",
			Year:   2024,
			Size:   200,
			Rate:   4.7,
		},
	}
	booksJSON, err := MarshalJSONSlice(originalBooks)
	if err != nil {
		t.Fatalf("Marshaling error: %v", err)
	}
	recoveredBooks, err := UnmarshalJSONSlice(booksJSON)
	if err != nil {
		t.Fatalf("Unmarshaling error: %v", err)
	}
	if !reflect.DeepEqual(originalBooks, recoveredBooks) {
		t.Errorf("Expected %v, received %v", originalBooks, recoveredBooks)
	}
}

func TestUnmarshalJSONSliceInvalidJSON(t *testing.T) {
	invalidJSON := `{
		 "id": "one",
		 "title": 2
	}`
	_, err := UnmarshalJSONSlice([]byte(invalidJSON))
	if err == nil {
		t.Errorf("An error was expected when unmarshaling incorrect JSON")
	}
}

func TestMarshalUnmarshalBook(t *testing.T) {
	originalBook := Book{
		ID:     1,
		Title:  "Test Book",
		Author: "Author",
		Year:   2024,
		Size:   100,
		Rate:   4.5,
	}
	bookJSON, err := json.Marshal(originalBook)
	if err != nil {
		t.Fatalf("Book marshaling error: %v", err)
	}

	var recoveredBook Book
	err = json.Unmarshal(bookJSON, &recoveredBook)
	if err != nil {
		t.Fatalf("Book unmarshaling error: %v", err)
	}

	if !reflect.DeepEqual(originalBook, recoveredBook) {
		t.Errorf("Expected %v, received %v", originalBook, recoveredBook)
	}
}

func TestMarshalJSONSliceEmpty(t *testing.T) {
	books := []Book{}
	booksJSON, err := MarshalJSONSlice(books)
	if err != nil {
		t.Errorf("Error when marshaling an empty slice: %v", err)
	}
	expected := "[]"
	if string(booksJSON) != expected {
		t.Errorf("Expected %v, received %v", expected, string(booksJSON))
	}
}

func TestUnmarshalJSONSliceNil(t *testing.T) {
	_, err := UnmarshalJSONSlice(nil)
	if err == nil {
		t.Errorf("An error was expected when unmarshaling nil")
	}
}

func TestUnmarshalJSONSliceEmptyString(t *testing.T) {
	_, err := UnmarshalJSONSlice([]byte(""))
	if err == nil {
		t.Errorf("An error was expected when unmarshaling an empty JSON string")
	}
}

func TestBookMarshalJSON(t *testing.T) {
	book := Book{
		ID:     1,
		Title:  "Test Book",
		Author: "Author",
		Year:   2024,
		Size:   100,
		Rate:   4.5,
	}

	expectedJSON := `{"id":1,"title":"Test Book","author":"Author","year":2024,"size":100,"rate":4.5}`
	bookJSON, err := book.MarshalJSON()
	if err != nil {
		t.Fatalf("Error in book marshaling: %v", err)
	}

	if string(bookJSON) != expectedJSON {
		t.Errorf("Expected %v, received %v", expectedJSON, string(bookJSON))
	}
}

func TestBookUnmarshalJSON(t *testing.T) {
	inputJSON := `{"id":1,"title":"Test Book","author":"Author","year":2024,"size":100,"rate":4.5}`
	var book Book
	err := book.UnmarshalJSON([]byte(inputJSON))
	if err != nil {
		t.Fatalf("Book demarshaling error: %v", err)
	}

	expectedBook := Book{
		ID:     1,
		Title:  "Test Book",
		Author: "Author",
		Year:   2024,
		Size:   100,
		Rate:   4.5,
	}

	if !reflect.DeepEqual(book, expectedBook) {
		t.Errorf("Expected %v, received %v", expectedBook, book)
	}
}

func TestBookUnmarshalJSONEmptyTitle(t *testing.T) {
	inputJSON := `{"id":1,"title":"","author":"Author","year":2024,"size":100,"rate":4.5}`
	var book Book
	err := book.UnmarshalJSON([]byte(inputJSON))
	if err == nil {
		t.Fatalf("An error was expected due to a missing book title")
	}
}

func TestMarshalProtobufSlice(t *testing.T) {
	books := []Book{
		{
			ID:     1,
			Title:  "Test Book 1",
			Author: "Author 1",
			Year:   2023,
			Size:   100,
			Rate:   4.5,
		},
		{
			ID:     2,
			Title:  "Test Book 2",
			Author: "Author 2",
			Year:   2024,
			Size:   200,
			Rate:   4.7,
		},
	}

	booksProto, err := MarshalProtobufSlice(books)
	if err != nil {
		t.Fatalf("Protobuf marshaling error: %v", err)
	}

	var pbBooks bookpb.Books
	err = proto.Unmarshal(booksProto, &pbBooks)
	if err != nil {
		t.Fatalf("Failed to unmarshal books Protobuf: %v", err)
	}

	if len(pbBooks.Books) != len(books) {
		t.Fatalf("Expected %d books, got %d", len(books), len(pbBooks.Books))
	}

	for i, pbBook := range pbBooks.Books {
		if books[i].ID != int(pbBook.Id) ||
			books[i].Title != pbBook.Title ||
			books[i].Author != pbBook.Author ||
			books[i].Year != int(pbBook.Year) ||
			books[i].Size != int(pbBook.Size) ||
			books[i].Rate != pbBook.Rate {
			t.Errorf("Mismatch in book details at index %d", i)
		}
	}
}

func TestUnmarshalProtobufSlice(t *testing.T) {
	// Сериализованные в Protobuf данные, полученные из вашего вывода
	inputProto, err := hex.DecodeString(
		"0a270801120b5465737420426f6f6b20311a08417574686f72203120e70f287b310000000000001240" +
			"0a280802120b5465737420426f6f6b20321a08417574686f72203220e80f28c80331cdcccccccccc1240",
	)
	if err != nil {
		t.Fatalf("Failed to decode hex string: %v", err)
	}

	expectedBooks := []Book{
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

	books, err := UnmarshalProtobufSlice(inputProto)
	if err != nil {
		t.Fatalf("Protobuf unmarshaling error: %v", err)
	}

	if !reflect.DeepEqual(books, expectedBooks) {
		t.Errorf("Expected %+v, received %+v", expectedBooks, books)
	}
}
