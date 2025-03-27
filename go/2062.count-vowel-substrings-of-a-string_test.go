package main

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func Test_countVowelSubstrings(t *testing.T) {
  tests := []struct {
    s    string
    want int
  }{
    {s: "cuaieuouac", want: 7},
    {s: "eiou", want: 0},
    {s: "aeiou", want: 1},
    {s: "aeiouu", want: 2},
    {s: "unicornarihan", want: 0},
  }
  for _, tt := range tests {
    tt := tt
    t.Run(tt.s, func(t *testing.T) {
      assert.Equalf(t, tt.want, countVowelSubstrings(tt.s), "countVowelSubstrings(%v)", tt.s)
    })
  }
}
