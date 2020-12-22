package main

import "testing"

func Test_validIPAddress(t *testing.T) {
	const (
		neither = "Neither"
		v4      = "IPv4"
		v6      = "IPv6"
	)

	tests := []struct {
		name string
		IP   string
		want string
	}{
		{IP: "12..0.0", want: neither},
		{IP: "02001:0db8:85a3:0:0:8A2E:0370:7334", want: neither},
		{IP: "2001::85a3:0:0:8A2E:0370:7334", want: neither},
		{IP: "2001:0db8:85a3:0:0:8A2E:0370:7334:", want: neither},
		{IP: "2001:0db8:85a3:0:0:8A2E:0370:7334", want: v6},
		{IP: "12.1.0.0", want: v4},
		{IP: "12.01.12.12", want: neither},
		{IP: "01.12.12.12", want: neither},
		{IP: "12.0.12.12", want: v4},
		{IP: "12.12.12.12", want: v4},
		{IP: "12.12.12.12.12", want: neither},
		{IP: "", want: neither},
		{IP: "12.12.", want: neither},
		{IP: "12.12.", want: neither},
		{IP: ".12", want: neither},
	}
	for _, tt := range tests {
		t.Run(tt.IP, func(t *testing.T) {
			if got := validIPAddress(tt.IP); got != tt.want {
				t.Errorf("validIPAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
