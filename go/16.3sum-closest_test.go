package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_threeSumClosest(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{nums: []int{1, 2, 4, 8, 16, 32, 64, 128}, k: 82, want: 82},
		{nums: []int{0, 2, 1, -3}, k: 1, want: 0},
		{nums: []int{-1, 2, 1, -4}, k: 1, want: 2},
	}
	for _, tc := range tests {
		tt := tc
		t.Run(fmt.Sprintln(tt.nums, tt.k, tt.want), func(t *testing.T) {
			if got := threeSumClosest(tt.nums, tt.k); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_threeSumClosest(b *testing.B) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{nums: []int{1, 2, 4, 8, 16, 32, 64, 128}, k: 82, want: 82},
		{nums: []int{0, 2, 1, -3}, k: 1, want: 0},
		{nums: []int{-1, 2, 1, -4}, k: 1, want: 2},
	}
	for _, tc := range tests {
		tt := tc
		for i := 0; i < b.N; i++ {
			b.Run(fmt.Sprintln(tt.nums, tt.k, tt.want), func(b *testing.B) {
				x := threeSumClosest(tt.nums, tt.k)
				_, _ = fmt.Fprintf(ioutil.Discard, "%d", x)
			})
		}
	}
}
