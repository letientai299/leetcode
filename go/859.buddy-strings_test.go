package main

import "testing"

func Test_buddyStrings(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want bool
	}{
		{a: "aa", b: "aa", want: true},
		{a: "ab", b: "ba", want: true},
		{a: "ab", b: "ab", want: false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := buddyStrings(tt.a, tt.b); got != tt.want {
				t.Errorf("buddyStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
