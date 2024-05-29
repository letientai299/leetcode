package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_answerQueries(t *testing.T) {
	type args struct {
		nums    []int
		queries []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums:    []int{4, 5, 2, 1},
				queries: []int{3, 10, 21},
			},
			want: []int{2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, answerQueries(tt.args.nums, tt.args.queries), "answerQueries(%v, %v)", tt.args.nums, tt.args.queries)
		})
	}
}
