package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addBoldTag(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s:     "aaabbb",
				words: []string{"aa", "b"},
			},
			want: "<b>aaabbb</b>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, addBoldTag(tt.args.s, tt.args.words), "addBoldTag(%v, %v)", tt.args.s, tt.args.words)
		})
	}
}
