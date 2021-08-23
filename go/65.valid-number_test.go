package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isNumber(t *testing.T) {
	valids := []string{"2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"}
	for _, s := range valids {
		t.Run("valid: "+s, func(t *testing.T) {
			assert.True(t, isNumber(s))
		})
	}

	invalids := []string{"abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53", "e1"}
	for _, s := range invalids {
		t.Run("invalid:"+s, func(t *testing.T) {
			assert.False(t, isNumber(s))
		})
	}
}
