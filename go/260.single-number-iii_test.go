package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{nums: []int{1, 1, 2, 2, -1, 4}, want: []int{-1, 4}},
		{nums: []int{1, 1, 2, 2, 3, 4}, want: []int{3, 4}},
		{nums: []int{1, 0}, want: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.nums); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
