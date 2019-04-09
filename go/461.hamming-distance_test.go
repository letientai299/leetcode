package main

import (
	"fmt"
	"testing"
)

func Test_hammingDistance(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		want int
	}{
		{1, 4, 2},
		{5, 10, 4},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.x, tt.y,
		)
		t.Run(testName, func(t *testing.T) {
			got := hammingDistance(tt.x, tt.y)
			if got != tt.want {
				t.Errorf("hammingDistance(%v, %v) = %v, want %v", tt.x, tt.y, got, tt.want)
			}
		})
	}
}
