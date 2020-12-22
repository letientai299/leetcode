package main

import "testing"

func Test_minAreaFreeRect(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   float64
	}{
		{points: [][]int{{0, 3}, {1, 2}, {3, 1}, {1, 3}, {2, 1}}, want: 0},
		{points: [][]int{{1, 2}, {2, 1}, {1, 0}, {0, 1}}, want: 2.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAreaFreeRect(tt.points); got != tt.want {
				t.Errorf("minAreaFreeRect() = %v, want %v", got, tt.want)
			}
		})
	}
}
