package main

import (
	"fmt"
	"testing"
)

func Test_countBinarySubstrings(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"00000011", 2},
		{"00110011", 6},
		{"10101", 4},
		{"001", 1},
		{"0011", 2},
		{"01", 1},
		{"", 0},
		{"1", 0},
		{"11", 0},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.s,
		)
		t.Run(testName, func(t *testing.T) {
			got := countBinarySubstrings(tt.s)
			if got != tt.want {
				t.Errorf("countBinarySubstrings(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

