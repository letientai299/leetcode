package main

import (
	"fmt"
	"testing"
)

func Test_nextGreatestLetter(t *testing.T) {
	tests := []struct {
		letters []byte
		target  byte
		want    byte
	}{
		{
			letters: []byte{'e', 'e', 'e', 'e', 'e', 'e', 'n', 'n', 'n', 'n'},
			target:  'n',
			want:    'e',
		},

		{
			letters: []byte{'c', 'f', 'j'},
			target:  'd',
			want:    'f',
		},

		{
			letters: []byte{'c', 'f', 'j'},
			target:  'c',
			want:    'f',
		},

		{
			letters: []byte{'c', 'f', 'j'},
			target:  'a',
			want:    'c',
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.letters, tt.target,
		)
		t.Run(testName, func(t *testing.T) {
			got := nextGreatestLetter(tt.letters, tt.target)
			if got != tt.want {
				t.Errorf("nextGreatestLetter(%v, %v) = %c, want %c", tt.letters, tt.target, got, tt.want)
			}
		})
	}
}
