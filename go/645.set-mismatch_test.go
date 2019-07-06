package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findErrorNums(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{

		{
			nums: []int{3, 3, 1},
			want: []int{2, 3},
		},

		{
			nums: []int{3, 2, 2},
			want: []int{1, 2},
		},

		{
			nums: []int{2, 3, 2},
			want: []int{2, 1},
		},

		{
			nums: []int{1, 2, 2, 4},
			want: []int{2, 3},
		},

		{
			nums: []int{1, 1, 3, 4},
			want: []int{1, 2},
		},

		{
			nums: []int{1, 2, 3, 1},
			want: []int{1, 4},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			got := findErrorNums(tt.nums)
			assert.ElementsMatchf(t, tt.want, got, "want %v, got %v", tt.want, got)
		})
	}
}
