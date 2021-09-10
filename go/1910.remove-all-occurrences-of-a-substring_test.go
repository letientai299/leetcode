package main

import "testing"

func Test_removeOccurrences(t *testing.T) {
	tests := []struct {
		s    string
		part string
		want string
	}{
		{s: "daabcbaabcbc", part: "abc", want: "dab"}}
	for _, tt := range tests {
		t.Run(tt.s+":"+tt.part, func(t *testing.T) {
			if got := removeOccurrences(tt.s, tt.part); got != tt.want {
				t.Errorf("removeOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}
