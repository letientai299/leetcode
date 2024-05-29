package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countPairsOfConnectableServers(t *testing.T) {
	type args struct {
		edges       [][]int
		signalSpeed int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				edges: [][]int{
					{0, 1, 1}, {1, 2, 5}, {2, 3, 13}, {3, 4, 9}, {4, 5, 2},
				},
				signalSpeed: 1,
			},
			want: []int{0, 4, 6, 6, 4, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countPairsOfConnectableServers(tt.args.edges, tt.args.signalSpeed), "countPairsOfConnectableServers(%v, %v)", tt.args.edges, tt.args.signalSpeed)
		})
	}
}
