package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getRow(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{30, []int{1, 30, 435, 4060, 27405, 142506, 593775, 2035800, 5852925, 14307150, 30045015, 54627300, 86493225, 119759850, 145422675, 155117520, 145422675, 119759850, 86493225, 54627300, 30045015, 14307150, 5852925, 2035800, 593775, 142506, 27405, 4060, 435, 30, 1}},
		{0, []int{1}},
		{1, []int{1, 1}},
		{2, []int{1, 2, 1}},
		{3, []int{1, 3, 3, 1}},
		{4, []int{1, 4, 6, 4, 1}},
		{5, []int{1, 5, 10, 10, 5, 1}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := getRow(tt.n)
			assert.Equal(t, got, tt.want)
		})
	}
}
