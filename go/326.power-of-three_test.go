package main

import (
	"fmt"
	"testing"
)

func Test_isPowerOfThree(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{0, false},
		{45, false},
		{1, true},
		{9, true},
		{27, true},
		{81, true},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.n,
		)
		t.Run(testName, func(t *testing.T) {
			got := isPowerOfThree(tt.n)
			if got != tt.want {
				t.Errorf("isPowerOfThree(%v) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
