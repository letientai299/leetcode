package main

import (
	"fmt"
	"testing"
)

func Test_canWinNim(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{4, false},
		{5, true},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.n,
		)
		t.Run(testName, func(t *testing.T) {
			got := canWinNim(tt.n)
			if got != tt.want {
				t.Errorf("canWinNim(%v) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
