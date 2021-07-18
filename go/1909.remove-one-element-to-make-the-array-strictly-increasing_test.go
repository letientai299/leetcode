package main

import "testing"

func Test_canBeIncreasing(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{nums: []int{2, 3, 1, 2}, want: false},
		{nums: []int{1, 2, 10, 5, 7}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeIncreasing(tt.nums); got != tt.want {
				t.Errorf("canBeIncreasing() = %v, want %v", got, tt.want)
			}
		})
	}
}
