package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubsequence(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{33, -27, -9, -83, 48},
				k:    2,
			},
			want: []int{33, 48},
		},

		{
			args: args{
				nums: []int{-1, -2, 3, 4},
				k:    2,
			},
			want: []int{3, 4},
		},

		{
			args: args{
				nums: []int{2, 1, 3, 3},
				k:    2,
			},
			want: []int{3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxSubsequence(tt.args.nums, tt.args.k), "maxSubsequence(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
