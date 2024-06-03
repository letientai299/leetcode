package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numberOfPoints(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: [][]int{{3, 6}, {1, 5}, {4, 7}},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numberOfPoints(tt.args.nums), "numberOfPoints(%v)", tt.args.nums)
		})
	}
}
