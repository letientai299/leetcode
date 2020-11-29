package main

import "testing"

func Test_brokenCalc(t *testing.T) {
	tests := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{x: 3, y: 15, want: 5},
		{x: 5, y: 5, want: 0},
		{x: 5, y: 8, want: 2},
		{x: 100, y: 128, want: 37},
		{x: 130, y: 128, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := brokenCalc(tt.x, tt.y); got != tt.want {
				t.Errorf("brokenCalc() = %v, want %v", got, tt.want)
			}
		})
	}
}
