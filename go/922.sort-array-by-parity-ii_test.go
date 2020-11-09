package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortArrayByParityII(t *testing.T) {
	tests := []struct {
		a    []int
		want []int
	}{
		{
			a: []int{4, 1, 1, 0, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := sortArrayByParityII(tt.a)
			for i, x := range got {
				assert.True(t, i%2 == x%2, "wrong", i, x, tt.a)
			}
		})
	}
}
