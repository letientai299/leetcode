package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findRelativeRanks(t *testing.T) {
	tests := []struct {
		nums []int
		want []string
	}{
		{
			nums: []int{5, 4, 3, 2, 1},
			want: []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			if got := findRelativeRanks(nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRelativeRanks(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
