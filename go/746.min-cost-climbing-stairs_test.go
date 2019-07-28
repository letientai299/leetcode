package main

import (
	"fmt"
	"testing"
)

func Test_minCostClimbingStairs(t *testing.T) {
	tests := []struct {
		cost []int
		want int
	}{
		{cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, want: 6},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.cost,
		)
		t.Run(testName, func(t *testing.T) {
			got := minCostClimbingStairs(tt.cost)
			if got != tt.want {
				t.Errorf("minCostClimbingStairs(%v) = %v, want %v", tt.cost, got, tt.want)
			}
		})
	}
}
