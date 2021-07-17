package main

import "testing"

func Test_largestOddNumber(t *testing.T) {
	tests := []struct {
		num  string
		want string
	}{
		{num: "52", want: "5"},
	}
	for _, tt := range tests {
		t.Run(tt.num, func(t *testing.T) {
			if got := largestOddNumber(tt.num); got != tt.want {
				t.Errorf("largestOddNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
