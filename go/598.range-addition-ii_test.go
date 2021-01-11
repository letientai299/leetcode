package main

import (
	"fmt"
	"testing"
)

func Test_maxCount(t *testing.T) {
	tests := []struct {
		m    int
		n    int
		ops  [][]int
		want int
	}{

		{
			m: 3,
			n: 3,
			ops: [][]int{
				{2, 2},
				{3, 3},
			},
			want: 4,
		},

		{
			m: 4,
			n: 3,
			ops: [][]int{
				{2, 2},
				{3, 3},
			},
			want: 4,
		},

		{
			m: 4,
			n: 3,
			ops: [][]int{
				{2, 3},
				{3, 3},
			},
			want: 6,
		},

		{
			m: 4,
			n: 3,
			ops: [][]int{
				{2, 1},
				{3, 3},
			},
			want: 2,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v, %v",
			tt.m, tt.n, tt.ops,
		)
		t.Run(testName, func(t *testing.T) {
			got := maxCount(tt.m, tt.n, tt.ops)
			if got != tt.want {
				t.Errorf("maxCount(%v, %v, %v) = %v, want %v", tt.m, tt.n, tt.ops, got, tt.want)
			}
		})
	}
}
