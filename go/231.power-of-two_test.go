package main

import (
	"fmt"
	"testing"
)

func Test_isPowerOfTwo(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{0, false},
		{1 << 0, true},
		{1 << 1, true},
		{1 << 2, true},
		{1 << 3, true},
		{1 << 4, true},
		{1 << 5, true},
		{1 << 6, true},
		{1 << 7, true},
		{1 << 8, true},
		{1 << 9, true},
		{1 << 10, true},
		{1 << 11, true},
		{1 << 12, true},
		{1 << 13, true},
		{13, false},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.n,
		)
		t.Run(testName, func(t *testing.T) {
			got := isPowerOfTwo(tt.n)
			if got != tt.want {
				t.Errorf("isPowerOfTwo(%v) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
