package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{74, 88, 74, 14, 486, 88, 14, 74, 14, 88}, want: 486},
		{nums: []int{0, 1, 0, 1, 0, 1, 99}, want: 99},
		{nums: []int{2, 2, 3, 2}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}

	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("rand-%d", i), func(t *testing.T) {
			rand.Seed(time.Now().Unix())
			n := rand.Int() % 10
			arr := make([]int, 0, 3*n+1)
			for i := 0; i < n; i++ {
				r := rand.Int() % 100
				arr = append(arr, r, r, r)
			}

			r := rand.Int() % 1000
			arr = append(arr, r)
			rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })

			got := singleNumber(arr)
			assert.Equalf(t, got, r, "arr=%v, want=%v", arr, r)
		})
	}
}
