package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumSubtreeSize(t *testing.T) {
	type args struct {
		edges  [][]int
		colors []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				edges:  [][]int{{0, 1}, {1, 2}},
				colors: []int{1, 1, 2}},
			want: 1,
		},

		{
			args: args{
				edges:  [][]int{{0, 1}, {0, 2}, {0, 3}},
				colors: []int{1, 1, 2, 3}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maximumSubtreeSize(tt.args.edges, tt.args.colors), "maximumSubtreeSize(%v, %v)", tt.args.edges, tt.args.colors)
		})
	}
}
