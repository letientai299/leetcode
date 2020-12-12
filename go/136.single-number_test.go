package main

import (
	"fmt"
	"testing"
)

func Test_singleNumber136(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.nums), func(t *testing.T) {
			if got := singleNumber136(tt.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
