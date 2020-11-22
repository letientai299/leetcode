package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combine(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want [][]int
	}{
		{n: 4, k: 3, want: [][]int{
			{1, 2, 3},
			{1, 2, 4},
			{1, 3, 4},
			{2, 3, 4},
		}},


		{n: 4, k: 1, want: [][]int{
			{1},
			{2},
			{3},
			{4},
		}},

		{n: 4, k: 2, want: [][]int{
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 3},
			{2, 4},
			{3, 4},
		}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := combine(tt.n, tt.k); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}
