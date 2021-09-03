package main

import "testing"

func Test_splitString(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{s: "0818079", want: true},
		{s: "108079", want: false},
		{s: "100", want: true},
		{s: "001", want: false},
		{s: "10009998", want: true},
		{s: "9080701", want: false},

		{s: "0050043", want: true},
		{s: "1234", want: false},
		{s: "4321", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := splitString(tt.s); got != tt.want {
				t.Errorf("splitString() = %v, want %v", got, tt.want)
			}
		})
	}
}
