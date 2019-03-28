package main

import (
	"fmt"
	"testing"
)

func Test_reverseString(t *testing.T) {
	tests := []struct {
		in  []byte
		out [] byte
	}{
		{
			in:  []byte{'h', 'e', 'l', 'l', 'o'},
			out: []byte{'o', 'l', 'l', 'e', 'h'},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.in,
		)
		t.Run(testName, func(t *testing.T) {
			reverseString(tt.in)
		})
	}
}
