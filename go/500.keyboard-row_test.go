package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	tests := []struct {
		words []string
		want  []string
	}{
		{
			words: []string{"Hello", "Alaska", "Dad", "Peace"},
			want:  []string{"Alaska", "Dad"},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.words,
		)
		t.Run(testName, func(t *testing.T) {
			if got := findWords(tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords(%v) = %v, want %v", tt.words, got, tt.want)
			}
		})
	}
}
