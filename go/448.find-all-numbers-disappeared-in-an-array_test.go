package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findDisappearedNumbers(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.nums,
		)
		t.Run(testName, func(t *testing.T) {
			if got := findDisappearedNumbers(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDisappearedNumbers(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
