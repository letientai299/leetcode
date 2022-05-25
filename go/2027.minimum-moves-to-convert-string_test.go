package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumMoves(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "XXX", want: 1},
		{s: "XXOX", want: 2},
		{s: "OOOO", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			assert.Equalf(t, tt.want, minimumMoves(tt.s), "minimumMoves(%v)", tt.s)
		})
	}
}
