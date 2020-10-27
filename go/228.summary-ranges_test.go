package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	tests := []struct {
		nums []int
		want []string
	}{
		{nums: []int{0, 1, 2, 4}, want: []string{"0->2", "4"}},
		{nums: []int{0, 1, 2}, want: []string{"0->2"}},
		{nums: []int{0}, want: []string{"0"}},
		{nums: []int{0, 2, 4}, want: []string{"0", "2", "4"}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.nums), func(t *testing.T) {
			if got := summaryRanges(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("summaryRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
