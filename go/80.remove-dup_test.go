package main

import "testing"

func Test_removeDuplicates80(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3}, want: 7},
		{nums: []int{2, 2, 2}, want: 2},
		{nums: []int{1, 1}, want: 2},
		{nums: []int{1, 2, 2}, want: 3},
		{nums: []int{1, 1, 2}, want: 3},
		{nums: []int{1}, want: 1},
		{nums: []int{}, want: 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := removeDuplicates80(tt.nums); got != tt.want {
				t.Errorf("removeDuplicates80() = %v, want %v, arr=%v", got, tt.want, tt.nums[:got])
			}
		})
	}
}
