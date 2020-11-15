package main

import (
	"reflect"
	"testing"
)

func Test_decrypt(t *testing.T) {
	tests := []struct {
		code []int
		k    int
		want []int
	}{
		{
			code: []int{2, 4, 9, 3},
			k:    1,
			want: []int{4, 9, 3, 2},
		},

		{
			code:
			[]int{2, 4, 9, 3},
			k:    -2,
			want: []int{12, 5, 6, 13},
		},

		{
			code:
			[]int{1, 2, 3, 4},
			k:    0,
			want: []int{0, 0, 0, 0},
		},

		{
			code:
			[]int{5, 7, 1, 4},
			k:    3,
			want: []int{12, 10, 16, 13},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := decrypt(tt.code, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
