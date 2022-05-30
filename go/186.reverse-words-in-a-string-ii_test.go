package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "the sky is blue",
			want: "blue is sky the",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := []byte(tt.name)
			reverseWords(bs)
			got := string(bs)
			assert.Equal(t, tt.want, got)
		})
	}
}
