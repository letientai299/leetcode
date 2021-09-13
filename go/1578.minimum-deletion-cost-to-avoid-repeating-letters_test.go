package main

import "testing"

func Test_minCost(t *testing.T) {
	tests := []struct {
		s    string
		cost []int
		want int
	}{
		{s: "bbbaaa", cost: []int{4, 9, 3, 8, 8, 9}, want: 23},
		{s: "abaac", cost: []int{1, 2, 3, 4, 5}, want: 3},
		{s: "abc", cost: []int{1, 2, 3}, want: 0},
		{s: "aabaa", cost: []int{1, 2, 3, 4, 1}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := minCost(tt.s, tt.cost); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
