package main

import (
	"reflect"
	"testing"
)

func Test_findRightInterval(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      []int
	}{
		{
			intervals: [][]int{{1, 4}, {2, 3}, {3, 4}},
			want:      []int{-1, 2, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRightInterval(tt.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRightInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
