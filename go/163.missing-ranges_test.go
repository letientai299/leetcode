package main

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_findMissingRanges(t *testing.T) {
	tests := []struct {
		nums  []int
		lower int
		upper int
		want  []string
	}{
		{
			nums: []int{-1}, lower: -2, upper: -1, want: []string{"-2"},
		},
		{
			nums: []int{2}, lower: 0, upper: 5, want: []string{"0->1", "3->5"},
		},
		{
			nums: []int{5}, lower: 0, upper: 5, want: []string{"0->4"},
		},
		{
			nums: []int{0}, lower: 0, upper: 5, want: []string{"1->5"},
		},
		{
			nums: []int{0, 1, 2}, lower: 0, upper: 5, want: []string{"3->5"},
		},
	}
	for i, tc := range tests {
		tt := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := findMissingRanges(tt.nums, tt.lower, tt.upper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMissingRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
