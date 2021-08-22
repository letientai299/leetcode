package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_diStringMatch(t *testing.T) {
	tests := []struct {
		s string
	}{
		{s: "III"},
		{s: "IDID"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := diStringMatch(tt.s)
			n := len(tt.s) + 1
			if !assert.Len(t, got, n) {
				return
			}

			for i := 0; i < n-1; i++ {
				a, b := got[i], got[i+1]
				if a < b {
					assert.EqualValues(t, tt.s[i], 'I')
				} else {
					assert.EqualValues(t, tt.s[i], 'D')
				}
			}
		})
	}
}
