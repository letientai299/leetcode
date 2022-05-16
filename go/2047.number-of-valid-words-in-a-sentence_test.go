package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countValidWords(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "!this  1-s b8d!", want: 0},
		{s: "alice and  bob are playing stone-game10", want: 5},
		{s: "", want: 0},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.s, func(t *testing.T) {
			assert.Equalf(t, tt.want, countValidWords(tt.s), "countValidWords(%v)", tt.s)
		})
	}
}
