package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numberOfSubarrays(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{nums: []int{1, 1, 1, 1, 1}, k: 1, want: 5},
		{nums: []int{1, 1, 2, 1, 1}, k: 3, want: 2},
		{nums: []int{2, 4, 6}, k: 1, want: 0},
		{nums: []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, k: 2, want: 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numberOfSubarrays(tt.nums, tt.k), "numberOfSubarrays(%v, %v)", tt.nums, tt.k)
		})
	}
}
