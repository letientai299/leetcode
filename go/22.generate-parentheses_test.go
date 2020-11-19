package main

import (
	"fmt"
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

func Test_generateParenthesis_2(t *testing.T) {
	for i := 0; i < 20; i++ {
		res := generateParenthesis(i)
		fmt.Printf("%10d %10d\n", i, len(res))
	}
}
