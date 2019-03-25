package main

import (
	"fmt"
	"testing"
)

func Test_countPrimes(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 0},
		{2, 0},
		{3, 1},
		{4, 2},
		{10, 4},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.n,
		)
		t.Run(testName, func(t *testing.T) {
			got := countPrimes(tt.n)
			if got != tt.want {
				t.Errorf("countPrimes(%v) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
