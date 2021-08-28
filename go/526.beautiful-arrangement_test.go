package main

import (
	"fmt"
	"testing"
)

func Test_countArrangement(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{n: 2, want: 2},
		{n: 1, want: 1},
		{n: 3, want: 3},
		{n: 4, want: 8},
		{n: 5, want: 10},
		{n: 6, want: 36},
		{n: 7, want: 41},
		{n: 8, want: 132},
		{n: 9, want: 250},
		{n: 10, want: 700},
		{n: 15, want: 24679},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.n, tt.want), func(t *testing.T) {
			if got := countArrangement(tt.n); got != tt.want {
				t.Errorf("countArrangement() = %v, want %v", got, tt.want)
			}
		})
	}
}
