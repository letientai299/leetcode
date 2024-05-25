package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_commonFactors(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{a: 6, b: 12, want: 4},
		{a: 12, b: 12, want: 6},
		{a: 25, b: 30, want: 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equalf(t, tt.want, commonFactors(tt.a, tt.b), "commonFactors(%v, %v)", tt.a, tt.b)
		})
	}
}
