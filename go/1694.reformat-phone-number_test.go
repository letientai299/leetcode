package main

import "testing"

func Test_reformatNumber(t *testing.T) {
	tests := []struct {
		number string
		want   string
	}{
		{number: "1-23-45 6", want: "123-456"},
		{number: "1-23-45 67", want: "123-45-67"},
		{number: "1-23-45 678", want: "123-456-78"},
		{number: "1-23-45 678-9", want: "123-456-789"},
	}
	for _, tt := range tests {
		t.Run(tt.number, func(t *testing.T) {
			if got := reformatNumber(tt.number); got != tt.want {
				t.Errorf("reformatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
