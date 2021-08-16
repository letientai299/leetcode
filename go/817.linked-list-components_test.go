package main

import "testing"

func Test_numComponents(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		nums []int
		want int
	}{
		{
			head: newList(3, 4, 0, 2, 1),
			nums: []int{4},
			want: 1,
		},

		{
			head: newList(0, 1, 2, 3, 4),
			nums: []int{0, 2, 4},
			want: 3,
		},

		{
			head: newList(0, 1, 2, 3, 4),
			nums: []int{0, 1, 3, 4},
			want: 2,
		},

		{
			head: newList(0, 1, 2, 3),
			nums: []int{0, 1, 3},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numComponents(tt.head, tt.nums); got != tt.want {
				t.Errorf("numComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}
