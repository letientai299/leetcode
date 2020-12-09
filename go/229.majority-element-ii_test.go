package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_majorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{nums: []int{1, 2, 3, 4}, want: []int{}},
		{nums: []int{1, 2, 3}, want: []int{}},
		{nums: []int{2, 2}, want: []int{2}},
		{nums: []int{1, 2, 2, 2, 1, 2}, want: []int{2}},
		{nums: []int{1, 2, 2, 2, 1}, want: []int{1, 2}},
		{nums: []int{1, 2, 2, 1, 3}, want: []int{1, 2}},
		{nums: nil, want: nil},
		{nums: []int{1}, want: []int{1}},
		{nums: []int{1, 2}, want: []int{1, 2}},
		{nums: []int{1, 2, 1}, want: []int{1}},
		{nums: []int{1, 2, 2, 1}, want: []int{1, 2}},
		{nums: []int{1, 2, 2, 0}, want: []int{2}},
		{nums: []int{2, 2, 2, 0}, want: []int{2}},
		{nums: []int{2, 1, 0, 1}, want: []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.nums); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
