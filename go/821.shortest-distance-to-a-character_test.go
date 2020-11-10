package main

import (
	"reflect"
	"testing"
)

func Test_shortestToChar(t *testing.T) {
	tests := []struct {
		s    string
		c    byte
		want []int
	}{
		{
			s: "loveleetcode", c: 'e',
			want: []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := shortestToChar(tt.s, tt.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestToChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
