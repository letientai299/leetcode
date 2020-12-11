package main

import (
	"reflect"
	"testing"
)

func Test_countBits(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want []int
	}{
		{num: 6, want: []int{0, 1, 1, 2, 1, 2, 2}},
		{num: 1, want: []int{0, 1}},
		{num: 2, want: []int{0, 1, 1}},
		{num: 3, want: []int{0, 1, 1, 2}},
		{num: 4, want: []int{0, 1, 1, 2, 1}},
		{num: 5, want: []int{0, 1, 1, 2, 1, 2}},
		{num: 10, want: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBits(tt.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
