package main

import (
	"reflect"
	"testing"
)

func Test_kWeakestRows(t *testing.T) {
	tests := []struct {
		mat  [][]int
		k    int
		want []int
	}{
		{
			mat: [][]int{
				{1, 0},
				{1, 0},
				{1, 0},
				{1, 1},
			},
			k:    4,
			want: []int{0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := kWeakestRows(tt.mat, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kWeakestRows() = %v, want %v", got, tt.want)
			}
		})
	}
}
