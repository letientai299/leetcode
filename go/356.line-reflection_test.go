package main

import "testing"

func Test_isReflected(t *testing.T) {
	tests := []struct {
		points [][]int
		want   bool
	}{
		{
			points: [][]int{{1, 2}, {2, 2}, {1, 4}, {2, 4}},
			want:   true,
		},

		{
			points: [][]int{{0, 0}, {0, -1}},
			want:   true,
		},

		{
			points: [][]int{{0, 0}, {0, 0}},
			want:   true,
		},

		{
			points: [][]int{{-16, 1}, {16, 1}, {16, 1}},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := isReflected(tt.points); got != tt.want {
				t.Errorf("isReflected() = %v, want %v", got, tt.want)
			}
		})
	}
}
