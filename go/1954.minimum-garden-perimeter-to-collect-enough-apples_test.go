package main

import "testing"

func Test_minimumPerimeter(t *testing.T) {
	tests := []struct {
		name   string
		apples int64
		want   int64
	}{
		{apples: 1000000000, want: 5040},
		{apples: 1, want: 8},
		{apples: 13, want: 16},
		{apples: 53, want: 16},
		{apples: 60, want: 16},
		{apples: 61, want: 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumPerimeter(tt.apples); got != tt.want {
				t.Errorf("minimumPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
