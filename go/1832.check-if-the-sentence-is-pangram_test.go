package main

import "testing"

func Test_checkIfPangram(t *testing.T) {
	tests := []struct {
		sentence string
		want     bool
	}{
		{sentence: "sdfsa", want: false},
		{sentence: "thequickbrownfoxjumpsoverthelazydog", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.sentence, func(t *testing.T) {
			if got := checkIfPangram(tt.sentence); got != tt.want {
				t.Errorf("checkIfPangram() = %v, want %v", got, tt.want)
			}
		})
	}
}
