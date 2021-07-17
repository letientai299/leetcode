package main

import "testing"

func Test_checkZeroOnes(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{name: "1101", want: true},
		{name: "110100010", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkZeroOnes(tt.name); got != tt.want {
				t.Errorf("checkZeroOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
