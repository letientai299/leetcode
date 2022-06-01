package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generatePalindromes(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{s: "aabb", want: []string{"abba", "baab"}},
		{s: "aabbb", want: []string{"abbba", "babab"}},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			assert.Equalf(t, tt.want, generatePalindromes(tt.s), "generatePalindromes(%v)", tt.s)
		})
	}
}
