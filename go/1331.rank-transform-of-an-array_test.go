package main

import (
	"reflect"
	"testing"
)

func Test_arrayRankTransform(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{4, 3, 2, 1},
			want: []int{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := arrayRankTransform(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arrayRankTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}
