package main

import "testing"

func Test_findPoisonedDuration(t *testing.T) {
	tests := []struct {
		name       string
		timeSeries []int
		duration   int
		want       int
	}{
		{timeSeries: []int{1, 2, 3, 4}, duration: 2, want: 5},
		{timeSeries: []int{1, 4}, duration: 2, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPoisonedDuration(tt.timeSeries, tt.duration); got != tt.want {
				t.Errorf("findPoisonedDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
