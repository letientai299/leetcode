package main

import (
	"testing"
)

func Test_mostCommonWord(t *testing.T) {
	tests := []struct {
		paragraph string
		banned    []string
		want      string
	}{
		{
			paragraph: "Bob",
			banned:    []string{},
			want:      "bob",
		},
		{
			paragraph: "Bob hit a ball, the hit BALL flew far after it was hit.",
			banned:    []string{"hit"},
			want:      "ball",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := mostCommonWord(tt.paragraph, tt.banned); got != tt.want {
				t.Errorf("mostCommonWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
