package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{3, 4, 5, 6},
			want: []int{120, 90, 12 * 6, 60},
		},

		{
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},

		{
			nums: []int{1, 2, 3},
			want: []int{6, 3, 2},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			if got := productExceptSelf(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
