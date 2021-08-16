package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rearrangeArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{nums: []int{1, 9, 2, 7, 12}},
		{nums: []int{1, 2, 5, 9}},
		{nums: []int{1, 2, 3}},
		{nums: []int{1, 2, 3, 4, 5}},
		{nums: []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rearrangeArray(tt.nums)
			t.Log(got)
			for i := 1; i < len(tt.nums)-1; i++ {
				mid := got[i]
				ave := (got[i+1] + got[i-1]) / 2
				assert.NotEqual(t, mid, ave, got)
			}
		})
	}
}
