package main

import (
	"reflect"
	"testing"
)

func Test_addToArrayForm(t *testing.T) {
	tests := []struct {
		a    []int
		k    int
		want []int
	}{

		{
			a:    []int{1, 9},
			k:    1,
			want: []int{2, 0},
		},

		{
			a:    []int{1, 2},
			k:    1208,
			want: []int{1, 2, 2, 0},
		},

		{
			a:    []int{1, 2},
			k:    1200,
			want: []int{1, 2, 1, 2},
		},

		{
			a:    []int{1, 2, 0, 0},
			k:    12,
			want: []int{1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := addToArrayForm(tt.a, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addToArrayForm() = %v, want %v", got, tt.want)
			}
		})
	}
}
