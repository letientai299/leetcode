package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},

		{
			numbers: []int{4, 7, 4, 15},
			target:  8,
			want:    []int{1, 3},
		},

		{
			numbers: []int{5, 7, 4, 15},
			target:  8,
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.numbers), func(t *testing.T) {
			if got := twoSum(tt.numbers, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
