package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intersect(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{},
		{
			nums1: []int{1, 1, 2},
			nums2: []int{1, 2, 3},
			want:  []int{1, 2},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.nums1, tt.nums2,
		)
		t.Run(testName, func(t *testing.T) {
			got := intersect350(tt.nums1, tt.nums2)
			assert.ElementsMatch(t, got, tt.want, "intersect(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
		})
	}
}
