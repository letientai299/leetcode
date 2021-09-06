package main

import "testing"

func Test_reductionOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{1, 2, 3}, want: 3},
		{nums: []int{1, 1, 2, 2, 3}, want: 4},
		{nums: []int{1, 1, 1}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reductionOperations(tt.nums); got != tt.want {
				t.Errorf("reductionOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
