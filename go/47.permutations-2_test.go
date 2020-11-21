package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_permuteUnique(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{1, 1, 2},
			want: [][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := permuteUnique(tt.nums); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
