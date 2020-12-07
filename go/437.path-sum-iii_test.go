package main

import (
	"fmt"
	"testing"
)

func Test_pathSum_427(t *testing.T) {
	tests := []struct {
		treeLevelOrder []int
		sum            int
		want           int
	}{

		{
			treeLevelOrder: []int{1, NA, 2, NA, 3, NA, 4, NA, 5},
			sum:            3,
			want:           2,
		},

		{
			treeLevelOrder: []int{1, NA, -1, NA, 1, NA, -1, NA, 1},
			sum:            1,
			want:           6,
		},

		{
			treeLevelOrder: []int{1, NA, -1, NA, 1, NA, -1, NA, 1},
			sum:            0,
			want:           6,
		},

		{
			treeLevelOrder: []int{10, 5, -3},
			sum:            5,
			want:           1,
		},

		{
			treeLevelOrder: []int{10, 5, -3, 3, 2, NA, 11, 3, -2, NA, 1},
			sum:            8,
			want:           3,
		},

		{
			treeLevelOrder: []int{10, 5, -3, 3, 2, NA, 11, 3, -2, NA, 1},
			sum:            -1,
			want:           0,
		},

		{
			treeLevelOrder: []int{10, 5, -3, 3, 2, NA, 11, 3, -2, NA, 1},
			sum:            1,
			want:           2,
		},
	}
	for _, tc := range tests {
		tt := tc
		root := treeFromLevelOrder(tt.treeLevelOrder...)
		testName := fmt.Sprintf(
			"%v, %v",
			tt.treeLevelOrder, tt.sum,
		)
		t.Run(testName, func(t *testing.T) {
			got := pathSum_437(root, tt.sum)
			if got != tt.want {
				t.Errorf("pathSum(%v, %v) = %v, want %v", tt.treeLevelOrder, tt.sum, got, tt.want)
			}
		})
	}
}
