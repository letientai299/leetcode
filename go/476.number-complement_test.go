package main

import (
	"fmt"
	"testing"
)

func Test_findComplement(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{5, 2},
		{3, 0},
		{9, 6},
		{1, 0},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.num,
		)
		t.Run(testName, func(t *testing.T) {
			got := findComplement(tt.num)
			if got != tt.want {
				t.Errorf("findComplement(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
