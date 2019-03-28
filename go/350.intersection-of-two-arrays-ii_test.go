package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_intersect(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.nums1, tt.nums2,
		)
		t.Run(testName, func(t *testing.T) {
			if got := intersect(tt.nums1, tt.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}
