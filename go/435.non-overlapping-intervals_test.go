package main

import "testing"

func Test_eraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{
			intervals: [][]int{{1, 2}, {2, 3}},
			want:      0,
		},

		{
			intervals: [][]int{{1, 2}, {1, 2}, {1, 2}},
			want:      2,
		},

		{
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			want:      1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eraseOverlapIntervals(tt.intervals); got != tt.want {
				t.Errorf("eraseOverlapIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
