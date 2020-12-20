package main

import "testing"

func Test_maxResult(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{1, -1, -2, 4, -7, 3}, k: 1, want: -2,
		},

		{
			nums: []int{1, -1, -2, 4, -7, 3}, k: 2, want: 7,
		},

		{
			nums: []int{10, -5, -2, 4, 0, 3}, k: 3, want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxResult(tt.nums, tt.k); got != tt.want {
				t.Errorf("maxResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
