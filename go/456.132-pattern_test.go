package main

import "testing"

func Test_find132pattern(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{nums: []int{3, 5, 0, 3, 4}, want: true},
		{nums: []int{1, 0, 1, -4, -3}, want: false},
		{nums: []int{1, 4, 0, -1, -2, -3, -1, -2}, want: true},
		{nums: []int{3, 1, 4, 2}, want: true},
		{nums: []int{5, 4, 2, 2, 1}, want: false},
		{nums: []int{1, 7, 0, 4, -1, 5, 2}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find132pattern(tt.nums); got != tt.want {
				t.Errorf("find132pattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
