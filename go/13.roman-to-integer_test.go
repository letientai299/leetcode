package main

import "testing"

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"I", 1},
		{"II", 2},
		{"III", 3},
		{"V", 5},
		{"IV", 4},
		{"VI", 6},
		{"VII", 7},
		{"VIII", 8},
		{"IX", 9},
		{"X", 10},
		{"XI", 11},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"CMXCIV", 994},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := romanToInt(tt.in); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
