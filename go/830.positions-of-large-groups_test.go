package main

import (
	"reflect"
	"testing"
)

func Test_largeGroupPositions(t *testing.T) {
	tests := []struct {
		s    string
		want [][]int
	}{
		{s: "aaa", want: [][]int{{0, 2}}},
		{s: "aaabb", want: [][]int{{0, 2}}},
		{s: "aaabbb", want: [][]int{{0, 2}, {3, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := largeGroupPositions(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largeGroupPositions() = %v, want %v", got, tt.want)
			}
		})
	}
}
