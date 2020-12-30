package main

import "testing"

func Test_decodeString(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{s: "2[abc]3[cd]ef", want: "abcabccdcdcdef"},
		{s: "y3[a2[b]]x", want: "yabbabbabbx"},
		{s: "3[a2[b]]", want: "abbabbabb"},
		{s: "3[ab]", want: "ababab"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := decodeString(tt.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
