package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pivotInteger(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{n: 8}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, pivotInteger(tt.args.n), "pivotInteger(%v)", tt.args.n)
		})
	}
}
