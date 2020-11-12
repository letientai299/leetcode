package main

import "testing"

func Test_isBoomerang(t *testing.T) {
	tests := []struct {
		points [][]int
		want   bool
	}{
		{
			points: [][]int{
				{0, 0},
				{1, 0},
				{2, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := isBoomerang(tt.points); got != tt.want {
				t.Errorf("isBoomerang() = %v, want %v", got, tt.want)
			}
		})
	}
}
