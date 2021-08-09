package main

import "testing"

func Test_maxWidthOfVerticalArea(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			points: [][]int{
				{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxWidthOfVerticalArea(tt.points); got != tt.want {
				t.Errorf("maxWidthOfVerticalArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
