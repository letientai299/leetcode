package main

import (
	"reflect"
	"testing"
)

func Test_distributeCandies(t *testing.T) {
	tests := []struct {
		n    int
		p    int
		want []int
	}{
		{n: 30, p: 2, want: []int{16, 14}},
		{n: 20, p: 2, want: []int{9, 11}},
		{n: 10, p: 2, want: []int{4, 6}},
		{n: 7, p: 2, want: []int{4, 3}},
		{n: 7, p: 4, want: []int{1, 2, 3, 1}},
		{n: 10, p: 3, want: []int{5, 2, 3}},
		{n: 7, p: 3, want: []int{2, 2, 3}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := distributeCandies(tt.n, tt.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
