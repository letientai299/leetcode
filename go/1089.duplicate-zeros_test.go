package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_duplicateZeros(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{0, 2, 0, 0, 1, 2, 3, 4},
			want: []int{0, 0, 2, 0, 0, 0, 0, 1},
		},

		{
			arr:  []int{1, 0, 2, 3, 0, 4, 5, 0},
			want: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			duplicateZeros(tt.arr)
			assert.Equal(t, tt.arr, tt.want)
		})
	}
}
