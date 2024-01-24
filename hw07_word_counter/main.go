package main

import (
	"fmt"
	"strings"
	"unicode"
)

func cleanup(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsNumber(r)
}

func countWords(text string) map[string]int {
	words := strings.FieldsFunc(text, cleanup)
	wordCount := make(map[string]int)

	for _, word := range words {
		cleanedWord := strings.ToLower(word)
		wordCount[cleanedWord]++
	}

	return wordCount
}

func main() {
	text := "Hello, world! This is a test. Hello,is world!"
	wordCount := countWords(text)

	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
