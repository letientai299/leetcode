package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeElement(t *testing.T) {
	tests := []struct {
		nums []int
		val  int
		want int
	}{
		{
			nums: []int{3, 2, 2, 3},
			val:  3,
			want: 2,
		},
		{
			nums: []int{},
			val:  3,
			want: 0,
		},
		{
			nums: []int{1, 1, 1, 1, 1, 1},
			val:  1,
			want: 0,
		},
		{
			nums: []int{1, 2, 3, 4, 5},
			val:  4,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d", tt.nums, tt.val), func(t *testing.T) {
			got := removeElement(tt.nums, tt.val)

			if got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}

			fmt.Println(tt.nums[:tt.want])
			for i := 0; i < tt.want; i++ {
				assert.NotEqual(t, tt.nums[i], tt.val)
			}
		})
	}
}
