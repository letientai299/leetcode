package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_construct2DArray(t *testing.T) {
	type args struct {
		original []int
		m        int
		n        int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{original: []int{1, 2, 3}, m: 2, n: 2},
			want: nil,
		},

		{
			args: args{original: []int{1, 2, 3, 4}, m: 2, n: 2},
			want: [][]int{
				{1, 2},
				{3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, construct2DArray(tt.args.original, tt.args.m, tt.args.n), "construct2DArray(%v, %v, %v)", tt.args.original, tt.args.m, tt.args.n)
		})
	}
}
