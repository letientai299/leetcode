package main

import "testing"

func Test_countNicePairs(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			nums: []int{42, 11, 1, 97},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNicePairs(tt.nums); got != tt.want {
				t.Errorf("countNicePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
