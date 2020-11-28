package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateParenthesis(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{
			n:    3,
			want: []string{"((()))", "(())()", "()(())", "()()()", "(()())"},
		},

		{
			n:    2,
			want: []string{"(())", "()()"},
		},

		{
			n:    1,
			want: []string{"()"},
		},

		{
			n:    0,
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			if got := generateParenthesis(tt.n); !assert.ElementsMatch(t, tt.want, got) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
