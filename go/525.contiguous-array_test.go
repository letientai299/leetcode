package main

import "testing"

func Test_findMaxLength(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{0, 1, 0, 1, 1, 1, 0, 0}, want: 8},
		{nums: []int{0, 1}, want: 2},
		{nums: []int{1, 0}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxLength(tt.nums); got != tt.want {
				t.Errorf("findMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
