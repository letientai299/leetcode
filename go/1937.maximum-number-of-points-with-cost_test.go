package main

import "testing"

func Test_maxPoints(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int64
	}{
		{
			points: [][]int{
				{1, 5}, {2, 3}, {4, 2},
			},
			want: 11,
		},

		{
			points: [][]int{
				{1, 2, 3}, {1, 5, 1}, {3, 1, 1},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPoints(tt.points); got != tt.want {
				t.Errorf("maxPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
