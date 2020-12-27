package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wiggleSort(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{nums: []int{4, 5, 5, 6}},
		{nums: []int{1, 2, 3}},
		{nums: []int{3, 2, 1}},
		{nums: []int{2, 1}},
		{nums: []int{1, 2}},
		{nums: []int{1, 3, 2, 2, 3, 1}},
		{nums: []int{1, 1, 1, 5, 6, 4}},
		{nums: []int{2, 2, 1, 3, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.nums)
			wiggleSort(tt.nums)
			up := true
			for i := 0; i < len(tt.nums)-1; i++ {
				if !assert.NotEqual(t, tt.nums[i], tt.nums[i+1]) {
					t.Logf("failed, nums=%v", tt.nums)
					return
				}

				if !assert.Equal(t, up, tt.nums[i] < tt.nums[i+1]) {
					t.Logf("failed, nums=%v", tt.nums)
					return
				}

				up = !up
			}
			t.Logf("ok, nums=%v", tt.nums)
		})
	}
}
