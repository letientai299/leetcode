package main

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElements(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			nums: []int{100, 1,   11,   1, 120, 111, 123,   1,  -1, -100},
			want: []int{120, 11, 120, 120, 123, 123,  -1, 100, 100,  100},
		},

		{nums: []int{6, 5, 4, 2, 1}, want: []int{-1, 6, 6, 6, 6}},
		{nums: []int{1, 1, 1, 1}, want: []int{-1, -1, -1, -1}},
		{nums: []int{1}, want: []int{-1}},
		{
			nums: []int{1, 2, 3, 3, 2, 1},
			want: []int{2, 3, -1, -1, 3, 2},
		},

		{
			nums: []int{1, 2, 1},
			want: []int{2, -1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElements(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
