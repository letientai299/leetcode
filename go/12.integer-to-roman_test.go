package main

import "testing"

func Test_intToRoman(t *testing.T) {
	tests := []struct {
		want string
		num  int
	}{
		{"MCMXCIV", 1994},
		{"V", 5},
		{"CMXCIV", 994},
		{"IX", 9},
		{"I", 1},
		{"II", 2},
		{"III", 3},
		{"IV", 4},
		{"VI", 6},
		{"VII", 7},
		{"VIII", 8},
		{"X", 10},
		{"XI", 11},
		{"LVIII", 58},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := intToRoman(tt.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
