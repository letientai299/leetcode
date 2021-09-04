package main

import "testing"

func Test_partitionDisjoint(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{1, 1}, want: 1},
		{nums: []int{5, 0, 3, 8, 6}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionDisjoint(tt.nums); got != tt.want {
				t.Errorf("partitionDisjoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
