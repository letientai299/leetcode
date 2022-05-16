package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_areNumbersAscending(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{s: "1 box has 3 blue 4 red 6 green and 12 yellow marbles", want: true},
		{s: "hello is 5 x 5", want: false},
		{s: "hello is 5 x 6", want: true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.s, func(t *testing.T) {
			assert.Equalf(t, tt.want, areNumbersAscending(tt.s), "areNumbersAscending(%v)", tt.s)
		})
	}
}
