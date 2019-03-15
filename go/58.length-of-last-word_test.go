package main

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"", 0},
		{"hello world", 5},
		{"Hello woRld", 5},
		{"Tiến Tài", 3},
		{"TiếnTài", 7},
		{"        TiếnTài", 7},
		{"TiếnTài        ", 7},
		{"anything else", 4},
		{"anything else             ", 4},
		{"else", 4},
		{"              else              ", 4},
		{"              else              a", 1},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := lengthOfLastWord(tt.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
