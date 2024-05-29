package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_alternateDigitSum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{n: 12}, want: -1},
		{args: args{n: 123}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, alternateDigitSum(tt.args.n), "alternateDigitSum(%v)", tt.args.n)
		})
	}
}
