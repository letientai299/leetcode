package main

import (
	"fmt"
	"testing"
)

func Test_rotatedDigits(t *testing.T) {
	tests := []struct {
		N    int
		want int
	}{
		{N: 10, want: 4},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.N,
		)
		t.Run(testName, func(t *testing.T) {
			got := rotatedDigits(tt.N)
			if got != tt.want {
				t.Errorf("rotatedDigits(%v) = %v, want %v", tt.N, got, tt.want)
			}
		})
	}
}
