package main

import (
	"fmt"
	"testing"
)

func Test_isPowerOfFour(t *testing.T) {
	tests := []struct {
		num  int
		want bool
	}{
		{1 << 0, true},
		{1 << 2, true},
		{1 << 4, true},
		{1 << 6, true},
		{1 << 8, true},
		{0, false},
		{2, false},
		{3, false},
		{5, false},
		{6, false},
		{7, false},
		{8, false},
		{9, false},
		{10, false},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.num,
		)
		t.Run(testName, func(t *testing.T) {
			got := isPowerOfFour(tt.num)
			if got != tt.want {
				t.Errorf("isPowerOfFour(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
