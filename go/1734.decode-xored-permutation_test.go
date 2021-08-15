package main

import (
	"reflect"
	"testing"
)

func Test_decode(t *testing.T) {
	tests := []struct {
		name    string
		encoded []int
		want    []int
	}{
		{
			encoded: []int{7, 5, 6, 11, 14, 15, 11, 6},
			want:    []int{6, 1, 4, 2, 9, 7, 8, 3, 5},
		},

		{
			encoded: []int{6, 5, 4, 6},
			want:    []int{2, 4, 1, 5, 3},
		},

		{
			encoded: []int{3, 1},
			want:    []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decode(tt.encoded); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
