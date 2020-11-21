package main

import "testing"

func Test_canJump(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{nums: []int{3, 0, 8, 2, 0, 0, 1}, want: true},
		{nums: []int{1, 1, 0, 1, 2}, want: false},
		{nums: []int{0}, want: true},
		{nums: []int{1, 2, 2, 6, 3, 6, 1, 8, 9, 4, 7, 6, 5, 6, 8, 2, 6, 1, 3, 6, 6, 6, 3, 2, 4, 9, 4, 5, 9, 8, 2, 2, 1, 6, 1, 6, 2, 2, 6, 1, 8, 6, 8, 3, 2, 8, 5, 8, 0, 1, 4, 8, 7, 9, 0, 3, 9, 4, 8, 0, 2, 2, 5, 5, 8, 6, 3, 1, 0, 2, 4, 9, 8, 4, 4, 2, 3, 2, 2, 5, 5, 9, 3, 2, 8, 5, 8, 9, 1, 6, 2, 5, 9, 9, 3, 9, 7, 6, 0, 7, 8, 7, 8, 8, 3, 5, 0}, want: true},
		{nums: []int{1, 1, 2}, want: true},
		{nums: []int{0, 1, 2}, want: false},
		{nums: []int{1, 2, 3}, want: true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := canJump(tt.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
