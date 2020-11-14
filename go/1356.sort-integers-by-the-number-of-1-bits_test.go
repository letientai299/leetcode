package main

import (
	"reflect"
	"testing"
)

func Test_sortByBits(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{10, 100, 1000, 10000},
			want: []int{10, 100, 10000, 1000},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := sortByBits(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortByBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
