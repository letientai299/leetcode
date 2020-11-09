package main

import (
	"reflect"
	"testing"
)

func Test_intervalIntersection(t *testing.T) {
	tests := []struct {
		a    [][]int
		b    [][]int
		want [][]int
	}{
		{
			a:    [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
			b:    [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			want: [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := intervalIntersection(tt.a, tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intervalIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
