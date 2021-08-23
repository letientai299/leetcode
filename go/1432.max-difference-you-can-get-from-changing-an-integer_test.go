package main

import "testing"

func Test_maxDiff(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{num: 123456, want: 820000},
		{num: 555, want: 888},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDiff(tt.num); got != tt.want {
				t.Errorf("maxDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
