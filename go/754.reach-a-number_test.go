package main

import (
	"fmt"
	"testing"
)

func Test_reachNumber(t *testing.T) {
	tests := []struct {
		target int
		want   int
	}{
		{target: 0, want: 0},
		{target: 1, want: 1},
		{target: 2, want: 3},
		{target: 3, want: 2},

		{target: -1, want: 1},
		{target: -2, want: 3},
		{target: -3, want: 2},

		{target: 4, want: 3},
		{target: 5, want: 5},
		{target: 6, want: 3},
		{target: 7, want: 5},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.target,
		)
		t.Run(testName, func(t *testing.T) {
			got := reachNumber(tt.target)
			if got != tt.want {
				t.Errorf("reachNumber(%v) = %v, want %v", tt.target, got, tt.want)
			}
		})
	}
}
