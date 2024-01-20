package main

import (
	"testing"
)

func TestGenerateChessboard(t *testing.T) {
	tests := []struct {
		name   string
		rows   int
		cols   int
		expect string
	}{

		{"empty board", 0, 0, ""},

		{"negative rows", -1, 8, ""},
		{"negative cols", 8, -1, ""},
		{"negative rows and cols", -1, -1, ""},

		{"2x2 board", 2, 2, " ------\n|    # |\n| #    |\n ------"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateChessboard(tt.rows, tt.cols); got != tt.expect {
				t.Errorf("generateChessboard(%d, %d) = %v, want %v", tt.rows, tt.cols, got, tt.expect)
			}
		})
	}
}
