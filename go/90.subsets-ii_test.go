package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsetsWithDup(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			nums: []int{1, 2, 2, 2, 3},
			want: [][]int{
				{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2},
				{1, 2, 2, 2}, {2, 2, 2}, {3}, {1, 3},
				{1, 2, 3}, {1, 2, 2, 3}, {1, 2, 2, 2, 3},
				{2, 3}, {2, 2, 3}, {2, 2, 2, 3},
			},
		},
		{
			nums: []int{1, 2, 2},
			want: [][]int{
				{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsetsWithDup(tt.nums); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("subsetsWithDup() = %v, want %v", got, tt.want)
			}
		})
	}
}
