package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumDifference(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{7, 1, 5, 4}, want: 4},
		{nums: []int{9, 4, 3, 2}, want: -1},
		{nums: []int{1, 5, 2, 10}, want: 9},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprint(tt.nums), func(t *testing.T) {
			assert.Equalf(t, tt.want, maximumDifference(tt.nums), "maximumDifference(%v)", tt.nums)
		})
	}
}
