package main

import (
	"fmt"
	"testing"
)

func Test_numberOfBoomerangs(t *testing.T) {
	tests := []struct {
		points [][]int
		want   int
	}{
		{
			points: [][]int{
				{0, 0}, {1, 0}, {2, 0},
			},
			want: 2,
		},

		{
			points: [][]int{
				{1, 1}, {2, 2}, {3, 3},
			},
			want: 2,
		},

		{
			points: [][]int{
				{0, 0}, {1, 0}, {-1, 0}, {0, 1}, {0, -1},
			},
			want: 20,
		},
	}
	for i, tc := range tests {
		if i != 1 {
			continue
		}
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.points,
		)
		t.Run(testName, func(t *testing.T) {
			got := numberOfBoomerangs(tt.points)
			if got != tt.want {
				t.Errorf("numberOfBoomerangs(%v) = %v, want %v", tt.points, got, tt.want)
			}
		})
	}
}
