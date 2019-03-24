package main

import (
	"fmt"
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	tests := []struct {
		num  uint32
		want int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{5, 2},
		{9, 2},
		{343, 6},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.num,
		)
		t.Run(testName, func(t *testing.T) {
			got := hammingWeight(tt.num)
			if got != tt.want {
				t.Errorf("hammingWeight(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
