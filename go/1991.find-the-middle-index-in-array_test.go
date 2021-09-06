package main

import "testing"

func Test_findMiddleIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{2, 3, -1, 8, 4}, want: 3},
		{nums: []int{0, 0, 0}, want: 0},
		{nums: []int{1, -1, 0}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMiddleIndex(tt.nums); got != tt.want {
				t.Errorf("findMiddleIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
