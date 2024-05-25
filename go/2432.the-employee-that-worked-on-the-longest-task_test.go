package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hardestWorker(t *testing.T) {
	type args struct {
		n    int
		logs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: 70,
				logs: [][]int{
					{36, 3}, {1, 5}, {12, 8}, {25, 9}, {53, 11}, {29, 12}, {52, 14},
				},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, hardestWorker(tt.args.n, tt.args.logs), "hardestWorker(%v, %v)", tt.args.n, tt.args.logs)
		})
	}
}
