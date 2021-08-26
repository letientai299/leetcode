package main

import (
	"testing"
)

func Test_validateBinaryTreeNodes(t *testing.T) {
	type args struct {
		n      int
		lefts  []int
		rights []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				n:      4,
				lefts:  []int{3, -1, 1, -1},
				rights: []int{-1, -1, 0, -1},
			},
			want: true,
		},

		{
			args: args{
				n:      4,
				lefts:  []int{1, 2, 0, -1},
				rights: []int{-1, -1, -1, -1},
			},
			want: false,
		},

		{
			args: args{
				n:      2,
				lefts:  []int{1, 0},
				rights: []int{-1, -1},
			},
			want: false,
		},

		{
			args: args{
				n:      4,
				lefts:  []int{1, -1, 3, -1},
				rights: []int{2, 3, -1, -1},
			},
			want: false,
		},

		{
			args: args{
				n:      4,
				lefts:  []int{1, -1, 3, -1},
				rights: []int{2, -1, -1, -1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateBinaryTreeNodes(tt.args.n, tt.args.lefts, tt.args.rights); got != tt.want {
				t.Errorf("validateBinaryTreeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
