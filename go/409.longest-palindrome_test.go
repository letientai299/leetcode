package main

import (
	"fmt"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"aaabbbcc", 7},
		{"abccccdd", 7},
		{"a", 1},
		{"", 0},
		{"Aa", 1},
		{"aabbb", 5},
		{"aabb", 4},
		{"Aabb", 3},
		{"Aabbbb", 5},
		{"AabbbB", 3},
		{"AabbbA", 5},
		{"civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth", 983},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.s,
		)
		t.Run(testName, func(t *testing.T) {
			got := longestPalindrome(tt.s)
			if got != tt.want {
				t.Errorf("longestPalindrome(%v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
