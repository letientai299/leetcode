package main

import (
	"reflect"
	"testing"
)

func Test_constructDistancedSequence(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{n: 5, want: []int{5, 3, 1, 4, 3, 5, 2, 4, 2}},
		{n: 1, want: []int{1}},
		{n: 2, want: []int{2, 1, 2}},
		{n: 3, want: []int{3, 1, 2, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructDistancedSequence(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructDistancedSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
