package main

import (
	"fmt"
	"testing"
)

func Test_largestTriangleArea(t *testing.T) {
	tests := []struct {
		points [][]int
		want   float64
	}{
		{
			points: [][]int{
				{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0},
			},
			want: 2,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.points,
		)
		t.Run(testName, func(t *testing.T) {
			got := largestTriangleArea(tt.points)
			if got != tt.want {
				t.Errorf("largestTriangleArea(%v) = %v, want %v", tt.points, got, tt.want)
			}
		})
	}
}
