package main

import (
	"reflect"
	"testing"
)

func Test_gardenNoAdj(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		paths [][]int
		want  []int
	}{
		{
			n:     4,
			paths: [][]int{{1, 2}, {3, 4}},
			want:  []int{1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gardenNoAdj(tt.n, tt.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gardenNoAdj() = %v, want %v", got, tt.want)
			}
		})
	}
}
