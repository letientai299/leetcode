package main

import "testing"

func Test_waysToMakeFair(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{1, 1, 1}, want: 3},
		{nums: []int{1, 2, 3}, want: 0},
		{nums: []int{2, 1, 6, 4}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToMakeFair(tt.nums); got != tt.want {
				t.Errorf("waysToMakeFair() = %v, want %v", got, tt.want)
			}
		})
	}
}
