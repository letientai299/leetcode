package main

import "testing"

func Test_minMoves2(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 3, 2}, want: 2},
		{nums: []int{0, 0, 3, 4, 6, 8}, want: 15},
		{nums: []int{0, 0, 3, 6, 8}, want: 14},
		{nums: []int{1, 2, 3, 6, 8}, want: 11},
		{nums: []int{1, 2, 3}, want: 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minMoves2(tt.nums); got != tt.want {
				t.Errorf("minMoves2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quickSelect(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{nums: []int{1, 3, 2}, k: 2, want: 2},
		{nums: []int{1, 2, 3, 4, 5}, k: 2, want: 2},
		{nums: []int{1, 2, 3, 4, 5}, k: 4, want: 4},
		{nums: []int{1, 2, 3, 4, 5}, k: 5, want: 5},
		{nums: []int{1, 2, 3, 4, 5}, k: 1, want: 1},
		{nums: []int{1, 2, 3, 4, 5, 6}, k: 1, want: 1},
		{nums: []int{1, 2, 3, 4, 5, 6}, k: 3, want: 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := quickSelect(tt.nums, tt.k); got != tt.want {
				t.Errorf("quickSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}
