package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intersection349(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			nums1: nil,
			nums2: nil,
			want:  nil,
		},
		{
			nums1: []int{1},
			nums2: nil,
			want:  nil,
		},
		{
			nums1: []int{1, 1},
			nums2: []int{1, 2},
			want:  []int{1},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.nums1, tt.nums2,
		)
		t.Run(testName, func(t *testing.T) {
			got := intersection349(tt.nums1, tt.nums2)
			assert.ElementsMatchf(t, got, tt.want, "intersection(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
		})
	}
}
