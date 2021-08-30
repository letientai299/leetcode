package main

import "testing"

func Test_minIncrementForUnique(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{3, 2, 1, 2, 1, 7}, want: 6},
		{nums: []int{1, 2, 2}, want: 1},
		{nums: []int{1, 2, 3}, want: 0},
		{nums: []int{2, 2, 3}, want: 2},
		{nums: []int{2, 2, 6}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minIncrementForUnique(tt.nums); got != tt.want {
				t.Errorf("minIncrementForUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
