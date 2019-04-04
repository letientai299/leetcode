package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	tests := []struct {
		s    string
		p    string
		want []int
	}{
		{
			s:    "cbaebabacd",
			p:    "abc",
			want: []int{0, 6},
		},

		{
			s:    "abab",
			p:    "ab",
			want: []int{0, 1, 2},
		},

		{
			s:    "abab",
			p:    "a",
			want: []int{0, 2},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.s, tt.p,
		)
		t.Run(testName, func(t *testing.T) {
			if got := findAnagrams(tt.s, tt.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams(%v, %v) = %v, want %v", tt.s, tt.p, got, tt.want)
			}
		})
	}
}
