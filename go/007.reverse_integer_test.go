package main

import (
	"strconv"
	"testing"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{
			input: 123,
			want:  321,
		},

		{
			input: 1234567890,
			want:  987654321,
		},

		{
			input: 199999999999,
			want:  0,
		},

		{
			input: -1231,
			want:  -1321,
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.input), func(t *testing.T) {
			if got := reverse(tt.input); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
