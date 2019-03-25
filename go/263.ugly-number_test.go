package main

import (
	"fmt"
	"testing"
)

func Test_isUgly(t *testing.T) {
	tests := []struct {
		num  int
		want bool
	}{
		{1000, true},
		{6, true},
		{8, true},
		{9, true},
		{10, true},
		{11, false},
		{0, false},
		{-1, false},
		{12, true},
		{15, true},
		{16, true},
		{17, false},
		{-1000, false},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.num,
		)
		t.Run(testName, func(t *testing.T) {
			got := isUgly(tt.num)
			if got != tt.want {
				t.Errorf("isUgly(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
