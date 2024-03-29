package main

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{0, 1, 2}, want: 3},
		{nums: []int{2, -11, 5}, want: 1},
		{nums: []int{0, 1, 1}, want: 2},
		{nums: []int{1, 1, 1}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
