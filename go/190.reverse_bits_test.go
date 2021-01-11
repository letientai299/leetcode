package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseBits(t *testing.T) {
	tests := []struct {
		num  uint32
		want uint32
	}{
		{
			num:  0,
			want: 0,
		},
		{
			num:  1,
			want: 1 << 31,
		},
		{
			num:  1 << 10,
			want: 1 << 21,
		},
	}
	for _, tt := range tests {
		b := fmt.Sprintf("%032b", tt.num)
		t.Run(b, func(t *testing.T) {
			got := reverseBits(tt.num)
			res := fmt.Sprintf("%032b", got)
			assert.Equal(t, _reverseStr(b), res)
		})
	}
}

func _reverseStr(someString string) string {
	runeString := []rune(someString)
	var reverseString string
	for i := len(runeString) - 1; i >= 0; i-- {
		reverseString += string(runeString[i])
	}
	return reverseString
}
