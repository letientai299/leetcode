package main

import (
	"fmt"
	"testing"
)

func Test_judgeSquareSum(t *testing.T) {
	tests := []struct {
		c    int
		want bool
	}{
		{0, true},
		{1, true},
		{2, true},
		{3, false},
		{4, true},
		{5, true},
		{6, false},
		{7, false},
		{8, true},
		{9, true},
		{10, true},
		{11, false},
		{12, false},
		{13, true},
		{14, false},
		{15, false},
		{16, true},
		{17, true},
		{18, true},
		{19, false},
		{20, true},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.c,
		)
		t.Run(testName, func(t *testing.T) {
			got := judgeSquareSum(tt.c)
			if got != tt.want {
				t.Errorf("judgeSquareSum(%v) = %v, want %v", tt.c, got, tt.want)
			}
		})
	}
}
