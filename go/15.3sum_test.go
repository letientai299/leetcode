package main

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := threeSum(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
