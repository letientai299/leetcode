package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDigit(t *testing.T) {
	tests := []struct {
		number string
		digit  byte
		want   string
	}{
		{number: "551", digit: '5', want: "51"},
		{number: "52651", digit: '5', want: "5261"},
		{number: "1231", digit: '1', want: "231"},
		{number: "123", digit: '3', want: "12"},
	}
	for _, tt := range tests {
		t.Run(tt.number, func(t *testing.T) {
			assert.Equalf(t, tt.want, removeDigit(tt.number, tt.digit), "removeDigit(%v, %v)", tt.number, tt.digit)
		})
	}
}
