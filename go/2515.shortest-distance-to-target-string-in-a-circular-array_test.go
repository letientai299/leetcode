package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_closetTarget(t *testing.T) {
	type args struct {
		words      []string
		target     string
		startIndex int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				words:      []string{"odjrjznxpn", "cyulttuabe", "zqxkdoeszk", "yeewpgriok", "odjrjznxpn", "btqpvxpjzv", "ukyudladhk", "ukyudladhk", "odjrjznxpn", "yeewpgriok"},
				target:     "odjrjznxpn",
				startIndex: 5,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, closetTarget(tt.args.words, tt.args.target, tt.args.startIndex), "closetTarget(%v, %v, %v)", tt.args.words, tt.args.target, tt.args.startIndex)
		})
	}
}
