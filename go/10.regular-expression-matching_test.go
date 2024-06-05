package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{s: "acaabbaccbbacaabbbb", p: "a*.*b*.*a*aa*a*"}, want: false},
		{args: args{s: "abc", p: "a.*c"}, want: true},
		{args: args{s: "ab", p: ".*c"}, want: false},
		{args: args{s: "a", p: "ab*"}, want: true},
		{args: args{s: "aaa", p: ".a"}, want: false},
		{args: args{s: "aa", p: ".a"}, want: true},
		{args: args{s: "aa", p: "a*"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isMatch(tt.args.s, tt.args.p), "isMatch(%v, %v)", tt.args.s, tt.args.p)
		})
	}
}
