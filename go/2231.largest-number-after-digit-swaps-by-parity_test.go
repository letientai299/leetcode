package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestInteger(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{num: 1234, want: 3412},
		{num: 65875, want: 87655},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, largestInteger(tt.num), "largestInteger(%v)", tt.num)
		})
	}
}
