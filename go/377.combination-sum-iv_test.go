package main

import "testing"

func Test_combinationSum4(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{nums: []int{1, 2, 3}, target: 4, want: 7},
		{nums: []int{1, 2, 3}, target: 10, want: 274},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum4(tt.nums, tt.target); got != tt.want {
				t.Errorf("combinationSum4() = %v, want %v", got, tt.want)
			}
		})
	}
}
