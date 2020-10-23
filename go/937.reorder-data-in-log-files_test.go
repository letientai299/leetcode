package main

import (
	"reflect"
	"testing"
)

func Test_reorderLogFiles(t *testing.T) {
	tests := []struct {
		logs []string
		want []string
	}{
		{
			logs: []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo", "a2 act car"},
			want: []string{"a2 act car", "g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"},
		},

		{
			logs: []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"},
			want: []string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"},
		},
	}
	for _, tc := range tests {
		tt := tc
		t.Run("", func(t *testing.T) {
			if got := reorderLogFiles(tt.logs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderLogFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
