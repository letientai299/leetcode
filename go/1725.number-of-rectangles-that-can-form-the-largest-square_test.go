package main

import "testing"

func Test_countGoodRectangles(t *testing.T) {
	tests := []struct {
		name       string
		rectangles [][]int
		want       int
	}{
		{
			rectangles: [][]int{
				{5, 8}, {3, 9}, {5, 12}, {16, 5},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodRectangles(tt.rectangles); got != tt.want {
				t.Errorf("countGoodRectangles() = %v, want %v", got, tt.want)
			}
		})
	}
}
