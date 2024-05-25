package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countTime(t *testing.T) {
	tests := []struct {
		time string
		want int
	}{
		{time: "07:?3", want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.time, func(t *testing.T) {
			assert.Equalf(t, tt.want, countTime(tt.time), "countTime(%v)", tt.time)
		})
	}
}
