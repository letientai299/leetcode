package main

import (
	"reflect"
	"testing"
)

func Test_generateMatrix(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]int
	}{
		{
			n: 4,
			want: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},

		{
			n:    1,
			want: [][]int{{1}},
		},
		{
			n:    2,
			want: [][]int{{1, 2}, {4, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
