package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kthDistinct(t *testing.T) {
	tests := []struct {
		arr  []string
		k    int
		want string
	}{
		{
			arr:  []string{"d", "b", "c", "b", "c", "a"},
			k:    2,
			want: "a",
		},

		{
			arr:  []string{"aaa", "aa", "a"},
			k:    1,
			want: "aaa",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s: %d", tt.arr, tt.k), func(t *testing.T) {
			assert.Equalf(t, tt.want, kthDistinct(tt.arr, tt.k), "kthDistinct(%v, %v)", tt.arr, tt.k)
		})
	}
}
