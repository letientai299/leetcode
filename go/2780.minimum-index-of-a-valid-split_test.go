package main

import (
  "fmt"
  "testing"

  "github.com/stretchr/testify/assert"
)

func Test_minimumIndex(t *testing.T) {
  tests := []struct {
    nums []int
    want int
  }{
    {nums: []int{2, 2, 2}, want: 0},
    {nums: []int{1, 2, 2, 2}, want: 2},
    {nums: []int{2, 1, 3, 1, 1, 1, 7, 1, 2, 1}, want: 4},
    {nums: []int{3, 3, 3, 3, 7, 2, 2}, want: -1},
  }
  for _, tt := range tests {
    t.Run(fmt.Sprint(tt.nums), func(t *testing.T) {
      assert.Equalf(t, tt.want, minimumIndex(tt.nums), "minimumIndex(%v)", tt.nums)
    })
  }
}
