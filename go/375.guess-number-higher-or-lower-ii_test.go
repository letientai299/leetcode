package main

import "testing"

func Test_getMoneyAmount(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{n: 1, want: 0},
		{n: 2, want: 1},
		{n: 3, want: 2},
		{n: 4, want: 4},
		{n: 5, want: 6},
		{n: 6, want: 8},
		{n: 7, want: 10},
		{n: 8, want: 12},
		{n: 9, want: 14},
		{n: 10, want: 16},
		{n: 11, want: 18},
		{n: 12, want: 21},
		{n: 13, want: 24},
		{n: 14, want: 27},
		{n: 15, want: 30},
		{n: 16, want: 34},
		{n: 25, want: 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMoneyAmount(tt.n); got != tt.want {
				t.Errorf("getMoneyAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
