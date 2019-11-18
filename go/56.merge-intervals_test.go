package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      [][]int
	}{
		{
			intervals: [][]int{
				{1, 3}, {2, 6}, {4, 5}, {10, 18},
			},

			want: [][]int{
				{1, 6}, {10, 18},
			},
		},

		{
			intervals: [][]int{
				{1, 3}, {2, 6}, {8, 10}, {15, 18},
			},

			want: [][]int{
				{1, 6}, {8, 10}, {15, 18},
			},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.intervals,
		)
		t.Run(testName, func(t *testing.T) {
			if got := merge(tt.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge(%v) = %v, want %v", tt.intervals, got, tt.want)
			}
		})
	}
}
