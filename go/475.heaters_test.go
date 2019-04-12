package main

import (
	"fmt"
	"testing"
)

func Test_findRadius(t *testing.T) {
	tests := []struct {
		houses  []int
		heaters []int
		want    int
	}{
		{[]int{1, 2, 3}, []int{3}, 2},
		{[]int{1, 2, 3, 4}, []int{1, 4}, 1},
		{[]int{1}, []int{1}, 0},
		{[]int{1, 5}, []int{1, 5}, 0},
		{[]int{1, 6}, []int{1, 5}, 1},
		{[]int{1, 2, 3}, []int{2}, 1},
		{[]int{1, 2, 3}, []int{1}, 2},
		{[]int{1, 2, 3}, []int{0}, 3},
		{[]int{1, 2, 3}, []int{4}, 3},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.houses, tt.heaters,
		)
		t.Run(testName, func(t *testing.T) {
			got := findRadius(tt.houses, tt.heaters)
			if got != tt.want {
				t.Errorf("findRadius(%v, %v) = %v, want %v", tt.houses, tt.heaters, got, tt.want)
			}
		})
	}
}
