package main

import (
	"fmt"
	"testing"
)

func Test_bulbSwitch(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 99, want: 9},
		{n: 25, want: 5},
		{n: 23, want: 4},
		{n: 22, want: 4},
		{n: 21, want: 4},
		{n: 20, want: 4},
		{n: 19, want: 4},
		{n: 18, want: 4},
		{n: 17, want: 4},
		{n: 16, want: 4},
		{n: 15, want: 3},
		{n: 14, want: 3},
		{n: 13, want: 3},
		{n: 12, want: 3},
		{n: 11, want: 3},
		{n: 10, want: 3},
		{n: 9, want: 3},
		{n: 8, want: 2},
		{n: 7, want: 2},
		{n: 6, want: 2},
		{n: 3, want: 1},
		{n: 4, want: 2},
		{n: 5, want: 2},
		{n: 6, want: 2},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.n,
		)
		t.Run(testName, func(t *testing.T) {
			got := bulbSwitch(tt.n)
			if got != tt.want {
				t.Errorf("bulbSwitch(%v) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
