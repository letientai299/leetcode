package main

import (
	"testing"
)

func Test_clumsy(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{n: 0, want: 0},
		{n: 1, want: 1},
		{n: 2, want: 2},
		{n: 3, want: 6},
		{n: 4, want: 7},
		{n: 5, want: 7},
		{n: 6, want: 8},
		{n: 7, want: 6},
		{n: 8, want: 9},
		{n: 9, want: 11},
		{n: 10, want: 12},
		{n: 11, want: 10}, // 3 -> -1
		{n: 12, want: 13}, // 0 -> +1
		{n: 13, want: 15}, // 1 -> +2
		{n: 14, want: 16}, // 2 -> +2
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clumsy(tt.n); got != tt.want {
				t.Errorf("clumsy() = %v, want %v", got, tt.want)
			}
		})
	}
}
