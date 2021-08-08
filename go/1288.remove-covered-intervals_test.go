package main

import "testing"

func Test_removeCoveredIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{
			intervals: [][]int{
				{1, 2}, {1, 4}, {3, 4},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeCoveredIntervals(tt.intervals); got != tt.want {
				t.Errorf("removeCoveredIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
