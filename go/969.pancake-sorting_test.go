package main

import (
	"reflect"
	"testing"
)

func Test_pancakeSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			arr:  []int{1, 2, 3},
			want: nil,
		},

		{
			arr:  []int{3, 2, 4, 1},
			want: []int{3, 4, 2, 3, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pancakeSort(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pancakeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
