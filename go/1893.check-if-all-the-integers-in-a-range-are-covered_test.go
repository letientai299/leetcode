package main

import "testing"

func Test_isCovered(t *testing.T) {
	tests := []struct {
		name   string
		ranges [][]int
		left   int
		right  int
		want   bool
	}{
		{
			ranges: [][]int{
				{1, 2}, {3, 4}, {5, 6},
			},
			left:  2,
			right: 5,
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCovered(tt.ranges, tt.left, tt.right); got != tt.want {
				t.Errorf("isCovered() = %v, want %v", got, tt.want)
			}
		})
	}
}
