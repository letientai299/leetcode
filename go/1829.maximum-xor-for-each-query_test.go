package main

import (
	"reflect"
	"testing"
)

func Test_getMaximumXor(t *testing.T) {
	tests := []struct {
		name       string
		nums       []int
		maximumBit int
		want       []int
	}{
		{
			nums:       []int{2, 3, 4, 7},
			maximumBit: 3,
			want:       []int{5, 2, 6, 5},
		},

		{
			nums:       []int{0, 1, 1, 3},
			maximumBit: 2,
			want:       []int{0, 3, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximumXor(tt.nums, tt.maximumBit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMaximumXor() = %v, want %v", got, tt.want)
			}
		})
	}
}
