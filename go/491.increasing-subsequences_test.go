package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findSubsequences(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			nums: []int{4, 3, 2, 1},
			want: [][]int{},
		},

		{
			nums: []int{4, 6, 7, 7},
			want: [][]int{
				{4, 6}, {4, 7}, {4, 6, 7}, {4, 6, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}, {4, 7, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubsequences(tt.nums); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("findSubsequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
