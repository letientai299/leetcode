package main

import (
	"reflect"
	"testing"
)

func Test_advantageCount(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{
			a:    []int{7, 6, 5, 7, 1, 3},
			b:    []int{4, 4, 5, 5, 2, 3},
			want: []int{6, 7, 7, 1, 3, 5},
		},

		{
			a:    []int{12, 24, 8, 32},
			b:    []int{13, 25, 32, 11},
			want: []int{24, 32, 8, 12},
		},

		{
			a:    []int{2, 7, 11, 15},
			b:    []int{1, 10, 4, 11},
			want: []int{2, 11, 7, 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := advantageCount(tt.a, tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("advantageCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
