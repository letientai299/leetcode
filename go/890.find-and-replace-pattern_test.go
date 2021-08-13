package main

import (
	"reflect"
	"testing"
)

func Test_findAndReplacePattern(t *testing.T) {
	tests := []struct {
		name    string
		words   []string
		pattern string
		want    []string
	}{
		{
			words:   []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"},
			pattern: "abb",
			want:    []string{"mee", "aqq"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAndReplacePattern(tt.words, tt.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAndReplacePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
