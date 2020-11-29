package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_baseNeg2(t *testing.T) {
	n := 10
	for i := -n; i <= n; i++ {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			got := baseNeg2(i)
			parsed := parseBaseNeg2(got)
			assert.Equal(t, parsed, i, "baseNeg2(%d) = %v, parsed %v", i, got, parsed)
		})
	}
}

func parseBaseNeg2(s string) int {
	v := 0
	for _, b := range s {
		v *= -2
		if b == '1' {
			v += 1
		}
	}
	return v
}
