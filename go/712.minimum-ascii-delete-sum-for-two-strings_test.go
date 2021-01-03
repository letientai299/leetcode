package main

import (
	"fmt"
	"testing"
)

func Test_minimumDeleteSum(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want int
	}{
		{b: "a", a: "abcbabbc", want: 687},
		{a: "ab", b: "abc", want: 'c'},
		{a: "sea", b: "eat", want: 231},
		{a: "sea", b: "tea", want: 231},
		{a: "sea", b: "seat", want: 116},
		{a: "a", b: "abcbabbc", want: 687},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.a, tt.b), func(t *testing.T) {
			if got := minimumDeleteSum(tt.a, tt.b); got != tt.want {
				t.Errorf("minimumDeleteSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
