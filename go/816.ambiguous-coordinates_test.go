package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ambiguousCoordinates(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{
			s: "(0123)",
			want: []string{
				"(0, 1.23)", "(0, 12.3)", "(0, 123)", "(0.1, 2.3)", "(0.1, 23)", "(0.12, 3)",
			},
		},

		{
			s: "(123)",
			want: []string{
				"(1, 2.3)", "(1, 23)", "(1.2, 3)", "(12, 3)",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := ambiguousCoordinates(tt.s)
			assert.ElementsMatchf(t, tt.want, got, "want=%v\ngot=%v", tt.want, got)
		})
	}
}
