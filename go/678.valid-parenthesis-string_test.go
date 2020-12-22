package main

import "testing"

func Test_checkValidString(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{s: "(())((())()()(*)(*()(())())())()()((()())((()))(*", want: false},
		{s: "()", want: true},
		{s: "(*)", want: true},
		{s: "(*", want: true},
		{s: "(*))", want: true},
		{s: "(((*))", want: true},
		{s: ")(", want: false},
		{s: "*)(", want: false},
		{s: "*)(*", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := checkValidString(tt.s); got != tt.want {
				t.Errorf("checkValidString() = %v, want %v", got, tt.want)
			}
		})
	}
}
