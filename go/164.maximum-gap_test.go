package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumGap(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			nums: []int{1, 10000000},
			want: 9999999,
		},

		{
			nums: []int{170, 45, 75, 90, 802, 24, 2, 66},
			want: 632,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maximumGap(tt.nums), "maximumGap(%v)", tt.nums)
		})
	}
}
