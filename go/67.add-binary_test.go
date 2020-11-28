package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addBinary(t *testing.T) {
	tests := []struct {
		a int64
		b int64
		c int64
	}{
		{1, 2, 3},
		{1, 20, 21},
		{10, 20, 30},
		{255, 1, 256},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			as := strconv.FormatInt(tt.a, 2)
			bs := strconv.FormatInt(tt.b, 2)
			got := addBinary(as, bs)
			c, _ := strconv.ParseInt(got, 2, 64)
			assert.Equal(t, tt.c, c)
		})
	}
}
