package main

import "testing"

func Test_maxProduct152(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{2, -1, 1, 1}, want: 2},
		{nums: []int{2, -5, -2, -4, 3}, want: 24},
		{nums: []int{0, 2}, want: 2},
		{nums: []int{2, 3, -2, 4}, want: 6},
		{nums: []int{-3, -2, -1, -4}, want: 24},
		{nums: []int{-3, -2, 0, -1, -4}, want: 6},
		{nums: []int{-3, 0, 0}, want: 0},
		{nums: []int{-3}, want: -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct152(tt.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
