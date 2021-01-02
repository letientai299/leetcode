package main

import (
	"fmt"
	"testing"
)

func Test_singleNonDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 1, 2, 2, 3, 4, 4, 5, 5, 6, 6, 7, 7}, want: 3},
		{nums: []int{1}, want: 1},
		{nums: []int{1, 1, 2}, want: 2},
		{nums: []int{1, 2, 2}, want: 1},

		{
			nums: []int{1, 1, 2, 2, 3, 3, 4},
			want: 4,
		},

		{
			nums: []int{1, 1, 2, 3, 3},
			want: 2,
		},

		{
			nums: []int{1, 1, 2, 2, 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.nums), func(t *testing.T) {
			if got := singleNonDuplicate(tt.nums); got != tt.want {
				t.Errorf("singleNonDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
