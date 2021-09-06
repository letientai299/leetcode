package main

import "testing"

func Test_countQuadruplets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			nums: []int{28, 8, 49, 85, 37, 90, 20, 8},
			want: 1,
		},

		{
			nums: []int{1, 2, 3, 6},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countQuadruplets(tt.nums); got != tt.want {
				t.Errorf("countQuadruplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
