package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		n     int
	}{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			n:     3,
			nums2: []int{2, 3, 5},
		},
		{
			nums1: []int{1, 2, 3, 0, 0, 0, 0, 0},
			n:     5,
			nums2: []int{2, 3, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			merge(tt.nums1, len(tt.nums1)-tt.n, tt.nums2, tt.n)

			for i := range tt.nums1 {
				if i == len(tt.nums1)-1 {
					return
				}

				assert.True(t, tt.nums1[i] <= tt.nums1[i+1])
			}
		})
	}
}
