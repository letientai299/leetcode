package main

import "testing"

func Test_complexNumberMultiply(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want string
	}{
		{
			a:    "1+1i",
			b:    "1+1i",
			want: "0+2i",
		},

		{
			a:    "1+-1i",
			b:    "1+-1i",
			want: "0+-2i",
		},

		{
			a:    "-1+-1i",
			b:    "-1+-1i",
			want: "0+2i",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := complexNumberMultiply(tt.a, tt.b); got != tt.want {
				t.Errorf("complexNumberMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
