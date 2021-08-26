package main

import "testing"

func Test_maskPII(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{s: "AB@QQ.com", want: "a*****b@qq.com"},

		{s: "1(234)567-890", want: "***-***-7890"},
		{s: "1-(10)12345678", want: "+*-***-***-5678"},
		{s: "86-(10)12345678", want: "+**-***-***-5678"},
		{s: "248-(10)12345678", want: "+***-***-***-5678"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := maskPII(tt.s); got != tt.want {
				t.Errorf("maskPII() = %v, want %v", got, tt.want)
			}
		})
	}
}
