package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_digitCount(t *testing.T) {
	tests := []struct {
		num  string
		want bool
	}{
		{num: "1210", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.num, func(t *testing.T) {
			assert.Equalf(t, tt.want, digitCount(tt.num), "digitCount(%v)", tt.num)
		})
	}
}
