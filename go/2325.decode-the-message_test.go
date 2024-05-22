package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_decodeMessage(t *testing.T) {
	tests := []struct {
		key     string
		message string
		want    string
	}{
		{
			key:     "eljuxhpwnyrdgtqkviszcfmabo",
			message: "zwx hnfx lqantp mnoeius ycgk vcnjrdb",
			want:    "the five boxing wizards jump quickly",
		},
	}
	for _, tt := range tests {
		t.Run(tt.key+"|"+tt.message, func(t *testing.T) {
			assert.Equalf(t, tt.want, decodeMessage(tt.key, tt.message), "decodeMessage(%v, %v)", tt.key, tt.message)
		})
	}
}

func Test_aa(t *testing.T) {
	fmt.Println(2 ^ 5)
	fmt.Println(2 & 5)
}
