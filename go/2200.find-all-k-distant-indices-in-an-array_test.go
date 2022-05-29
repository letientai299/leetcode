package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findKDistantIndices(t *testing.T) {
	type args struct {
		nums []int
		key  int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{2, 1, 1, 1, 2},
				key:  2,
				k:    1,
			},
			want: []int{0, 1, 3, 4},
		},
	
		{
			args: args{
				nums: []int{3, 4, 9, 1, 3, 9, 5},
				k:    1,
				key:  9,
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findKDistantIndices(tt.args.nums, tt.args.key, tt.args.k), "findKDistantIndices(%v, %v, %v)", tt.args.nums, tt.args.key, tt.args.k)
		})
	}
}
