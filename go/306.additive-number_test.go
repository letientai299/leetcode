package main

import "testing"

func Test_isAdditiveNumber(t *testing.T) {
	tests := []struct {
		num  string
		want bool
	}{
		{num: "0235813", want: false},
		{num: "211738", want: true},
		{num: "000", want: true},
		{num: "101", want: true},
		{num: "112", want: true},
		{num: "1099100199", want: false},
		{num: "1123581321", want: true},
		{num: "199100199", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.num, func(t *testing.T) {
			if got := isAdditiveNumber(tt.num); got != tt.want {
				t.Errorf("isAdditiveNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
