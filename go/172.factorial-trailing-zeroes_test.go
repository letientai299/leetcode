package main

import (
	"fmt"
	"testing"
)

func Test_trailingZeroes(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{5, 1},
		{10, 2},
		{30, 7},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.n,
		)
		t.Run(testName, func(t *testing.T) {
			got := trailingZeroes(tt.n)
			if got != tt.want {
				t.Errorf("trailingZeroes(%v) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
