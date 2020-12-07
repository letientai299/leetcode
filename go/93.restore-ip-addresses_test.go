package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_restoreIpAddresses(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{
			s: "010010",
			want: []string{
				"0.10.0.10",
				"0.100.1.0",
			},
		},
		{
			s: "101023",
			want: []string{
				"1.0.10.23",
				"10.1.0.23",
				"101.0.2.3",
				"1.0.102.3",
				"10.10.2.3",
			},
		},

		{
			s: "25525522",
			want: []string{
				"2.55.255.22",
				"25.5.255.22",
				"25.52.55.22",
				"255.2.55.22",
				"255.25.5.22",
				"255.25.52.2",
				"255.255.2.2",
			},
		},
		{
			s: "0000", want: []string{"0.0.0.0"},
		},
		{
			s: "1111", want: []string{"1.1.1.1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := restoreIpAddresses(tt.s); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
