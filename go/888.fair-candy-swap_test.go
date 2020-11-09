package main

import (
	"reflect"
	"testing"
)

func Test_fairCandySwap(t *testing.T) {
	tests := []struct {
		a    []int
		b    []int
		want []int
	}{
		{a: []int{3, 5}, b: []int{4, 2}, want: []int{3, 2}},
		{a: []int{1, 1}, b: []int{2, 2}, want: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := fairCandySwap(tt.a, tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fairCandySwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
