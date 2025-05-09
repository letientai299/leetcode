package main

import (
  "fmt"
  "testing"

  "github.com/stretchr/testify/assert"
)

func Test_minimumPairRemoval(t *testing.T) {
  tests := []struct {
    nums []int
    want int
  }{
    {nums: []int{5, 2, 3, 1}, want: 2},
    {nums: []int{1, 3, 2}, want: 2},
    {nums: []int{1}, want: 0},
    {nums: []int{1, 2, 3}, want: 0},
  }
  for _, tt := range tests {
    t.Run(fmt.Sprint(tt.nums), func(t *testing.T) {
      assert.Equalf(t, tt.want, minimumPairRemoval(tt.nums), "minimumPairRemoval(%v)", tt.nums)
    })
  }
}
