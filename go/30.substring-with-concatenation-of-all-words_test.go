package main

import (
	"reflect"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	tests := []struct {
		s     string
		words []string
		want  []int
	}{
		{
			s:     "aaaaaaaaaaaaaa",
			words: []string{"aa", "aa"},
			want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},

		{
			s:     "foo",
			words: []string{"foo", "bar", "bar"},
			want:  nil,
		},

		{
			s:     "xfoobarbarxbarfoobar",
			words: []string{"foo", "bar", "bar"},
			want:  []int{1, 11},
		},

		{
			s:     "xbarfoobar",
			words: []string{"foo", "bar", "bar"},
			want:  []int{1},
		},

		{
			s:     "barfoobarbarfoo",
			words: []string{"foo", "bar", "bar"},
			want:  []int{0, 3, 6},
		},

		{
			s:     "barfoothefoobarman",
			words: []string{"foo", "bar"},
			want:  []int{0, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := findSubstring(tt.s, tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
