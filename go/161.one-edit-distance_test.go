package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isOneEditDistance(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{s: "", t: ""}, want: false},
		{args: args{s: "ab", t: "abc"}, want: true},
		{args: args{s: "ab", t: "acb"}, want: true},
		{args: args{s: "ab", t: "cab"}, want: true},
		{args: args{s: "abx", t: "cab"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isOneEditDistance(tt.args.s, tt.args.t), "isOneEditDistance(%v, %v)", tt.args.s, tt.args.t)
		})
	}
}
