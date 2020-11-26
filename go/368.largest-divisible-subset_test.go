package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestDivisibleSubset(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 2, 3, 8},
			want: []int{1, 2, 8},
		},

		{
			nums: []int{2, 3},
			want: []int{2},
		},

		{
			nums: []int{1, 2, 3},
			want: []int{1, 2},
		},

		{
			nums: []int{1, 2, 4, 8},
			want: []int{1, 2, 4, 8},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := largestDivisibleSubset(tt.nums); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("largestDivisibleSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}
