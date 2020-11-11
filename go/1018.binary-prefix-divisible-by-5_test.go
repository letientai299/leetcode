package main

import (
	"reflect"
	"testing"
)

func Test_prefixesDivBy5(t *testing.T) {
	tests := []struct {
		a    []int
		want []bool
	}{
		{
			a:    []int{1, 0, 1},
			want: []bool{false, false, true},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := prefixesDivBy5(tt.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prefixesDivBy5() = %v, want %v", got, tt.want)
			}
		})
	}
}
