package main

import "testing"

func Test_trap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{height: []int{3, 2, 1}, want: 0},

		{
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},

		{height: []int{1, 2, 1}, want: 0},
		{height: []int{1, 1, 1}, want: 0},
		{height: []int{1, 2, 3}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
