package main

import (
	"reflect"
	"testing"
)

func Test_indexPairs(t *testing.T) {
	tests := []struct {
		text  string
		words []string
		want  [][]int
	}{
		{
			text:  "thestoryofleetcodeandme",
			words: []string{"story", "leetcode", "me"},
			want:  [][]int{{3, 7}, {10, 17}, {21, 22}},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := indexPairs(tt.text, tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("indexPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
