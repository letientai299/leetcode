package main

import (
	"fmt"
	"testing"
)

func Test_canConstruct(t *testing.T) {
	tests := []struct {
		ransomNote string
		magazine   string
		want       bool
	}{
		{
			ransomNote: "a",
			magazine:   "b",
			want:       false,
		},

		{
			ransomNote: "a",
			magazine:   "ab",
			want:       true,
		},

		{
			ransomNote: "aa",
			magazine:   "baa",
			want:       true,
		},

		{
			ransomNote: "bg",
			magazine:   "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj",
			want:       true,
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.ransomNote, tt.magazine,
		)
		t.Run(testName, func(t *testing.T) {
			got := canConstruct(tt.ransomNote, tt.magazine)
			if got != tt.want {
				t.Errorf("canConstruct(%v, %v) = %v, want %v", tt.ransomNote, tt.magazine, got, tt.want)
			}
		})
	}
}
