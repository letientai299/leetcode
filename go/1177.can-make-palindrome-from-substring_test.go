package main

import (
	"reflect"
	"testing"
)

func Test_canMakePaliQueries(t *testing.T) {
	tests := []struct {
		s       string
		queries [][]int
		want    []bool
	}{
		{
			s:       "abcda",
			queries: [][]int{{3, 3, 0}, {1, 2, 0}, {0, 3, 1}, {0, 3, 2}, {0, 4, 1}},
			want:    []bool{true, false, false, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := canMakePaliQueries(tt.s, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canMakePaliQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
