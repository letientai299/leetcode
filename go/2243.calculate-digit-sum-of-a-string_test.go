package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_digitSum(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want string
	}{
		{s: "11111222223", k: 3, want: "135"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, digitSum(tt.s, tt.k), "digitSum(%v, %v)", tt.s, tt.k)
		})
	}
}
