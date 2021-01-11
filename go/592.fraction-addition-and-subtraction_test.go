package main

import "testing"

func Test_fractionAddition(t *testing.T) {
	tests := []struct {
		expression string
		want       string
	}{
		{expression: "-5/3+1/3", want: "-4/3"},
		{expression: "-5/3-1/3", want: "-2/1"},
		{expression: "-5/3+5/3", want: "0/1"},
		{expression: "-5/3+1/4", want: "-17/12"},
		{expression: "5/3-1/3", want: "4/3"},
		{expression: "5/3+1/3", want: "2/1"},
		{expression: "-1/2+1/2+3/3", want: "1/1"},
		{expression: "-1/2+1/2+1/3", want: "1/3"},
	}
	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			if got := fractionAddition(tt.expression); got != tt.want {
				t.Errorf("fractionAddition() = %v, want %v", got, tt.want)
			}
		})
	}
}
