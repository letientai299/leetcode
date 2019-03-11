package main

import (
	"strconv"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs []string
		want string
	}{
		{[]string{"aa", "a"}, "a"},
		{[]string{"flower", "flight", "floor"}, "fl"},
		{[]string{"flower", "flood", "floor"}, "flo"},
		{[]string{"flower", "flight", "xxxx"}, ""},
		{[]string{"dog", "cat", "dooo"}, ""},
		{[]string{"dog", "mog", "fdafsfafdogfdasf"}, ""},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := longestCommonPrefix(tt.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
