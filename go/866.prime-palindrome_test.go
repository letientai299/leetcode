package main

import (
	"fmt"
	"testing"
)

func Test_primePalindrome(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 12822, want: 13331},
		{n: 100040000, want: 100050001},
		{n: 10000000, want: 100030001},
		{n: 10000, want: 10301},
		{n: 102, want: 131},
		{n: 1, want: 2},
		{n: 2, want: 2},
		{n: 3, want: 3},
		{n: 4, want: 5},
		{n: 5, want: 5},
		{n: 6, want: 7},
		{n: 7, want: 7},
		{n: 8, want: 11},
		{n: 9, want: 11},
		{n: 10, want: 11},
		{n: 11, want: 11},
		{n: 12, want: 101},
		{n: 13, want: 101},
		{n: 14, want: 101},
		{n: 15, want: 101},
		{n: 16, want: 101},
		{n: 17, want: 101},
		{n: 18, want: 101},
		{n: 19, want: 101},
		{n: 20, want: 101},
		{n: 21, want: 101},
		{n: 22, want: 101},
		{n: 23, want: 101},
		{n: 24, want: 101},
		{n: 25, want: 101},
		{n: 26, want: 101},
		{n: 27, want: 101},
		{n: 28, want: 101},
		{n: 29, want: 101},
		{n: 30, want: 101},
		{n: 31, want: 101},
		{n: 32, want: 101},
		{n: 33, want: 101},
		{n: 34, want: 101},
		{n: 35, want: 101},
		{n: 36, want: 101},
		{n: 37, want: 101},
		{n: 38, want: 101},
		{n: 39, want: 101},
		{n: 40, want: 101},
		{n: 41, want: 101},
		{n: 42, want: 101},
		{n: 43, want: 101},
		{n: 44, want: 101},
		{n: 45, want: 101},
		{n: 46, want: 101},
		{n: 47, want: 101},
		{n: 48, want: 101},
		{n: 49, want: 101},
		{n: 50, want: 101},
		{n: 51, want: 101},
		{n: 52, want: 101},
		{n: 53, want: 101},
		{n: 54, want: 101},
		{n: 55, want: 101},
		{n: 56, want: 101},
		{n: 57, want: 101},
		{n: 58, want: 101},
		{n: 59, want: 101},
		{n: 60, want: 101},
		{n: 61, want: 101},
		{n: 62, want: 101},
		{n: 63, want: 101},
		{n: 64, want: 101},
		{n: 65, want: 101},
		{n: 66, want: 101},
		{n: 67, want: 101},
		{n: 68, want: 101},
		{n: 69, want: 101},
		{n: 70, want: 101},
		{n: 71, want: 101},
		{n: 72, want: 101},
		{n: 73, want: 101},
		{n: 74, want: 101},
		{n: 75, want: 101},
		{n: 76, want: 101},
		{n: 77, want: 101},
		{n: 78, want: 101},
		{n: 79, want: 101},
		{n: 80, want: 101},
		{n: 81, want: 101},
		{n: 82, want: 101},
		{n: 83, want: 101},
		{n: 84, want: 101},
		{n: 85, want: 101},
		{n: 86, want: 101},
		{n: 87, want: 101},
		{n: 88, want: 101},
		{n: 89, want: 101},
		{n: 90, want: 101},
		{n: 91, want: 101},
		{n: 92, want: 101},
		{n: 93, want: 101},
		{n: 94, want: 101},
		{n: 95, want: 101},
		{n: 96, want: 101},
		{n: 97, want: 101},
		{n: 98, want: 101},
		{n: 99, want: 101},
		{n: 100, want: 101},
		{n: 101, want: 101},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d=>%d", tt.n, tt.want), func(t *testing.T) {
			if got := primePalindrome(tt.n); got != tt.want {
				t.Errorf("primePalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextPotentialPalindromeNum(t *testing.T) {
	ds := digits(12822)
	for i := 0; i < 100; i++ {
		ds = nextPotentialPalindromeNum(ds)
		fmt.Println(ds)
	}
}
