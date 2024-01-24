package main

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected map[string]int
	}{
		{
			name: "simple sentence",
			text: "Hello, world! This is a test. Hello, world!",
			expected: map[string]int{
				"hello": 2,
				"world": 2,
				"this":  1,
				"is":    1,
				"a":     1,
				"test":  1,
			},
		},
		{
			name: "with numbers and punctuation",
			text: "A man, a plan, a canal: Panama 123 123!",
			expected: map[string]int{
				"a":      3,
				"man":    1,
				"plan":   1,
				"canal":  1,
				"panama": 1,
				"123":    2,
			},
		},
		{
			name: "case sensitivity",
			text: "Go gO GO",
			expected: map[string]int{
				"go": 3,
			},
		},
		{
			name:     "empty string",
			text:     "",
			expected: map[string]int{},
		},
		{
			name:     "spaces and tabs only",
			text:     "     \t\t\t   ",
			expected: map[string]int{},
		},
		{
			name:     "special characters",
			text:     "#$@!*&^%${}[]:;'",
			expected: map[string]int{},
		},
		{
			name: "mixed spaces and words",
			text: "   mixed   spaces   words  ",
			expected: map[string]int{
				"mixed":  1,
				"spaces": 1,
				"words":  1,
			},
		},
		{
			name: "russian characters",
			text: "Привет, мир! Привет, мир!",
			expected: map[string]int{
				"привет": 2,
				"мир":    2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countWords(tt.text)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Test '%s' failed: countWords(%q) = %v; want %v", tt.name, tt.text, got, tt.expected)
			}
		})
	}
}
