package main

import "testing"

func Test_distanceBetweenBusStops(t *testing.T) {
	tests := []struct {
		distance []int
		a        int
		b        int
		want     int
	}{
		{distance: []int{1, 2, 3, 4}, a: 0, b: 3, want: 4},
		{distance: []int{1, 2, 3, 4}, a: 0, b: 1, want: 1},
		{distance: []int{8, 11, 6, 7, 10, 11, 2}, a: 0, b: 3, want: 25},
		{distance: []int{1, 2, 3, 4}, a: 0, b: 2, want: 3},
		{distance: []int{1, 2, 3, 1}, a: 0, b: 2, want: 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := distanceBetweenBusStops(tt.distance, tt.a, tt.b); got != tt.want {
				t.Errorf("distanceBetweenBusStops() = %v, want %v", got, tt.want)
			}
		})
	}
}
