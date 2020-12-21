package main

import (
	"reflect"
	"testing"
)

func Test_pathInZigZagTree(t *testing.T) {
	tests := []struct {
		name  string
		label int
		want  []int
	}{
		{label: 32, want: []int{1, 2, 7, 8, 31, 32}},
		{label: 2, want: []int{1, 2}},
		{label: 26, want: []int{1, 2, 6, 10, 26}},
		{label: 14, want: []int{1, 3, 4, 14}},
		{label: 3, want: []int{1, 3}},
		{label: 1, want: []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathInZigZagTree(tt.label); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathInZigZagTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
