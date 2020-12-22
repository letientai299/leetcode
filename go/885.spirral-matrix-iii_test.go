package main

import (
	"reflect"
	"testing"
)

func Test_spiralMatrixIII(t *testing.T) {
	type args struct {
		R  int
		C  int
		r0 int
		c0 int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				R:  1,
				C:  4,
				r0: 0,
				c0: 0,
			},
			want: [][]int{
				{0, 0},
				{0, 1},
				{0, 2},
				{0, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralMatrixIII(tt.args.R, tt.args.C, tt.args.r0, tt.args.c0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralMatrixIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
