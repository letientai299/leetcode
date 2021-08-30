package main

import "testing"

func Test_minAreaRect(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			points: [][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {4, 1}, {4, 3}},
			want:   2,
		},
		{
			points: [][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2}},
			want:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAreaRect(tt.points); got != tt.want {
				t.Errorf("minAreaRect() = %v, want %v", got, tt.want)
			}
		})
	}
}
