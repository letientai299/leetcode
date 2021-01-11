package main

import "testing"

func Test_numTriplets(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			nums1: []int{1, 1},
			nums2: []int{1, 1, 1},
			want:  9,
		},

		{
			nums1: []int{4, 7, 9, 11, 23},
			nums2: []int{3, 5, 1024, 12, 18},
			want:  0,
		},

		{
			nums1: []int{7, 7, 8, 3},
			nums2: []int{1, 2, 9, 7},
			want:  2,
		},

		{
			nums1: []int{1, 1},
			nums2: []int{1},
			want:  1,
		},

		{
			nums1: []int{4, 7},
			nums2: []int{5, 2, 8, 9},
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTriplets(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("numTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
