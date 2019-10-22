package main

import (
	"fmt"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	tests := []struct {
		A    []int
		B    []int
		want float64
	}{
		{
			A:    []int{1, 5, 6, 8},
			B:    []int{2, 3, 4, 7},
			want: 4.5,
		},

		{
			A:    []int{1},
			B:    []int{2, 3},
			want: 2,
		},

		{
			A:    []int{7, 8},
			B:    []int{1, 2, 3, 4, 5, 6},
			want: 4.5,
		},

		{
			A:    []int{},
			B:    []int{1, 3},
			want: 2,
		},

		{
			A:    []int{2},
			B:    []int{1},
			want: 1.5,
		},

		{
			A:    []int{1, 2, 3, 4, 5, 6},
			B:    []int{7, 8},
			want: 4.5,
		},

		{
			A:    []int{2},
			B:    []int{},
			want: 2,
		},

		{
			A:    []int{},
			B:    []int{1},
			want: 1,
		},


		{
			A:    []int{1, 3},
			B:    []int{},
			want: 2,
		},

		{
			A:    []int{1, 3},
			B:    []int{2},
			want: 2.0,
		},

		{
			A:    []int{1, 3},
			B:    []int{2, 4},
			want: 2.5,
		},

		{
			A:    []int{1, 3, 4},
			B:    []int{2},
			want: 2.5,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.A, tt.B,
		)
		t.Run(testName, func(t *testing.T) {
			got := findMedianSortedArrays(tt.A, tt.B)
			if got != tt.want {
				t.Errorf("findMedianSortedArrays(%v, %v) = %v, want %v", tt.A, tt.B, got, tt.want)
			}
		})
	}
}
