package main

import (
	"fmt"
	"testing"
)

func Test_arrangeCoins(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 1, want: 1},
		{n: 2, want: 1},
		{n: 3, want: 2},
		{n: 4, want: 2},
		{n: 5, want: 2},
		{n: 6, want: 3},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.n,
		)
		t.Run(testName, func(t *testing.T) {
			got := arrangeCoins(tt.n)
			if got != tt.want {
				t.Errorf("arrangeCoins(%v) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
