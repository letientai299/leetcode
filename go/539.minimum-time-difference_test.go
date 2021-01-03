package main

import "testing"

func Test_findMinDifference(t *testing.T) {
	tests := []struct {
		name       string
		timePoints []string
		want       int
	}{
		{timePoints: []string{"12:12", "00:12", "12:13"}, want: 1},
		{timePoints: []string{"00:00", "04:00", "22:00"}, want: 120},
		{timePoints: []string{"23:59", "00:00"}, want: 1},
		{timePoints: []string{"23:59", "00:00", "00:00"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinDifference(tt.timePoints); got != tt.want {
				t.Errorf("findMinDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
