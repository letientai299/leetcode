package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_groupStrings(t *testing.T) {
	tests := []struct {
		name    string
		strings []string
		want    [][]string
	}{
		{
			strings: []string{"aa", "bb", "b"},
			want: [][]string{
				{"b"}, {"aa", "bb"},
			},
		},

		{
			strings: []string{"az", "ba", "a", "z"},
			want: [][]string{
				{"a", "z"}, {"az", "ba"},
			},
		},

		{
			strings: []string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"},
			want: [][]string{
				{"acef"}, {"a", "z"}, {"abc", "bcd", "xyz"}, {"az", "ba"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupStrings(tt.strings)
			want := tt.want
			for i, ws := range got {
				sort.Strings(ws)
				got[i] = ws
			}

			for i, ws := range want {
				sort.Strings(ws)
				want[i] = ws
			}

			assert.ElementsMatchf(t, want, got, "groupStrings(%v)", tt.strings)
		})
	}
}
