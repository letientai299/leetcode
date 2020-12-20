package main

import "testing"

func Test_stoneGameVII(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   int
	}{
		{
			stones: []int{5, 3, 1, 4, 2},
			want:   6,
		},

		{
			stones: []int{7, 90, 5, 1, 100, 10, 10, 2},
			want:   122,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGameVII(tt.stones); got != tt.want {
				t.Errorf("stoneGameVII() = %v, want %v", got, tt.want)
			}
		})
	}
}
