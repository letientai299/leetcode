package main

import "testing"

func Test_numDifferentIntegers(t *testing.T) {
	tests := []struct {
		word string
		want int
	}{
		{word: "0i00e", want: 1},
		{word: "a1b01c001", want: 1},
		{word: "192383183928778851682383a2089984061937879119", want: 2},
		{word: "a123bc34d8ef34", want: 3},
		{word: "leet1234code234", want: 2},
		{word: "leet1234code1234", want: 1},
		{word: "leet1234code0001234", want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			if got := numDifferentIntegers(tt.word); got != tt.want {
				t.Errorf("numDifferentIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
