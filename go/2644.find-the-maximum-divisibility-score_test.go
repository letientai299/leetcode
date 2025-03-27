package main

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func Test_maxDivScore(t *testing.T) {
  tests := []struct {
    nums     []int
    divisors []int
    want     int
  }{
    {
      nums:     []int{73, 13, 20, 6},
      divisors: []int{56, 75, 83, 26, 24, 53, 56, 61},
      want:     24,
    },
  }
  for _, tt := range tests {
    t.Run("", func(t *testing.T) {
      assert.Equalf(t, tt.want, maxDivScore(tt.nums, tt.divisors), "maxDivScore(%v, %v)", tt.nums, tt.divisors)
    })
  }
}
