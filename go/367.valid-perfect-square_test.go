package main

import (
	"fmt"
	"testing"
)

func Test_isPerfectSquare(t *testing.T) {
	tests := []struct {
		num  int
		want bool
	}{
		{1, true},
		{4, true},
		{9, true},
		{16, true},
		{25, true},
		{36, true},
		{81, true},
		{801, false},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.num,
		)
		t.Run(testName, func(t *testing.T) {
			got := isPerfectSquare(tt.num)
			if got != tt.want {
				t.Errorf("isPerfectSquare(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
