package main

import (
	"reflect"
	"testing"
)

func Test_restoreArray(t *testing.T) {
	tests := []struct {
		name          string
		adjacentPairs [][]int
		want          []int
	}{
		{
			adjacentPairs: [][]int{
				{2, 1},
				{3, 4},
				{3, 2},
			},
			want: []int{1,2,3,4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreArray(tt.adjacentPairs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
