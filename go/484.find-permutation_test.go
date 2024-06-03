package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findPermutation(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int
	}{
		{s: "I", want: []int{1, 2}},
		{s: "DDII", want: []int{3, 2, 1, 4, 5}},
		{s: "D", want: []int{2, 1}},
		{s: "DI", want: []int{2, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findPermutation(tt.s), "findPermutation(%v)", tt.s)
		})
	}
}
