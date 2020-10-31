package main

import (
	"fmt"
	"testing"
)

func Test_minCost(t *testing.T) {
	tests := []struct {
		costs [][]int
		want  int
	}{
		{
			costs: [][]int{
				{17, 2, 17}, {16, 16, 5}, {14, 3, 19},
			},
			want: 10,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.costs,
		)
		t.Run(testName, func(t *testing.T) {
			got := minCost(tt.costs)
			if got != tt.want {
				t.Errorf("minCost(%v) = %v, want %v", tt.costs, got, tt.want)
			}
		})
	}
}
