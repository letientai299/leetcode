package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_captureForts(t *testing.T) {
	type args struct {
		forts []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{forts: []int{1, 0, 0, -1, 0, 0, -1, 0, 0, 1}}, want: 2},
		{args: args{forts: []int{1, 0, 0, -1, 0, 0, 0, 0, 1}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, captureForts(tt.args.forts), "captureForts(%v)", tt.args.forts)
		})
	}
}
