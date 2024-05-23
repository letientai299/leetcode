package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_repeatedCharacter(t *testing.T) {
	tests := []struct {
		s    string
		want byte
	}{
		{s: "abccbaacz", want: 'c'},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			assert.Equalf(t, tt.want, repeatedCharacter(tt.s), "repeatedCharacter(%v)", tt.s)
		})
	}
}
