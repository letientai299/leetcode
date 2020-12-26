package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partition(t *testing.T) {
	tests := []struct {
		s    string
		want [][]string
	}{
		{
			s: "abbab",
			want: [][]string{
				{"a", "b", "b", "a", "b"}, {"a", "b", "bab"}, {"a", "bb", "a", "b"}, {"abba", "b"},
			},
		},

		{
			s:    "bb",
			want: [][]string{{"b", "b"}, {"bb"}},
		},

		{
			s: "aab",
			want: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := partition(tt.s); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
