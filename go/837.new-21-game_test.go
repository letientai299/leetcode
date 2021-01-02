package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_new21Game(t *testing.T) {
	type args struct {
		n int
		k int
		w int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{args: args{n: 9676, k: 8090, w: 3056}, want: 0.76813},
		{args: args{n: 21, k: 17, w: 10}, want: 0.73278},
		{args: args{n: 6, k: 1, w: 10}, want: 0.6},
		{args: args{n: 3, k: 3, w: 2}, want: 0.625},
		{args: args{n: 7, k: 7, w: 6}, want: 0.25360},
		{args: args{n: 5, k: 5, w: 4}, want: 0.36035},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := new21Game(tt.args.n, tt.args.k, tt.args.w)
			d := tt.want - got
			if d < 0 {
				d = -d
			}
			assert.LessOrEqualf(t, d, 1e-5, "want=%.5f, got=%.5f", tt.want, got)
		})
	}
}

// 0.25*(0.25+0.25
