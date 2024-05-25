package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_equalFrequency(t *testing.T) {
	tests := []struct {
		word string
		want bool
	}{
		{word: "zz", want: true},
		{word: "bbac", want: true},
		{word: "ddaccb", want: false},
		{word: "bac", want: true},
		{word: "bacc", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			assert.Equalf(t, tt.want, equalFrequency(tt.word), "equalFrequency(%v)", tt.word)
		})
	}
}
