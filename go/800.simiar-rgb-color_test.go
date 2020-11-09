package main

import "testing"

func Test_similarRGB(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{s: "#09f166", want: "#11ee66"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := similarRGB(tt.s); got != tt.want {
				t.Errorf("similarRGB() = %v, want %v", got, tt.want)
			}
		})
	}
}
