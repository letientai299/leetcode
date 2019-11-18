package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsets(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 2}, want: 1 << 2},
		{nums: []int{1}, want: 1 << 1},
		{nums: []int{1, 2, 3}, want: 1 << 3},
		{nums: []int{1, 2, 3, 4}, want: 1 << 4},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			got := subsets(tt.nums)
			fmt.Println(got)
			assert.Equal(t, len(got), tt.want)
		})
	}
}
