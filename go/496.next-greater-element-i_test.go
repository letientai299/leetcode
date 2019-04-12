package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_nextGreaterElement(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{

		{
			nums1: []int{4, 1, 2},
			nums2: []int{1, 3, 4, 2},
			want:  []int{-1, 3, -1},
		},

		{
			nums1: []int{2, 1, 3},
			nums2: []int{2, 3, 1},
			want:  []int{3, -1, -1},
		},

		{
			nums1: []int{1, 2},
			nums2: []int{1, 3, 4, 2},
			want:  []int{3, -1},
		},

		{
			nums1: []int{2, 4},
			nums2: []int{1, 2, 3, 4},
			want:  []int{3, -1},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.nums1, tt.nums2,
		)
		t.Run(testName, func(t *testing.T) {
			in := make([]int, len(tt.nums1))
			copy(in, tt.nums1)
			if got := nextGreaterElement(in, tt.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}
