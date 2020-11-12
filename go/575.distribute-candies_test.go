package main

import (
	"fmt"
	"testing"
)

func Test_distributeCandies_575(t *testing.T) {
	tests := []struct {
		candies []int
		want    int
	}{
		{candies: []int{1, 2}, want: 1},
		{candies: []int{1, 1}, want: 1},
		{candies: []int{1, 2, 1, 2}, want: 2},
		{candies: []int{1, 1, 1, 2, 1, 3, 1, 4}, want: 4},
		{candies: []int{1, 1, 3, 2}, want: 2},
		{candies: []int{1, 2, 3, 4}, want: 2},
		{candies: []int{1, 1, 1, 4}, want: 2},
		{candies: []int{1, 1, 1, 1}, want: 1},
		{candies: []int{1, 1, 2, 2, 3, 3}, want: 3},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.candies,
		)
		t.Run(testName, func(t *testing.T) {
			got := distributeCandies_575(tt.candies)
			if got != tt.want {
				t.Errorf("distributeCandies(%v) = %v, want %v", tt.candies, got, tt.want)
			}
		})
	}
}
