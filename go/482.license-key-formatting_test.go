package main

import (
	"fmt"
	"testing"
)

func Test_licenseKeyFormatting(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want string
	}{
		{
			s:    "ab",
			k:    2,
			want: "AB",
		},
		{
			s:    "ab",
			k:    1,
			want: "A-B",
		},
		{
			s:    "55F3Z2E9W",
			k:    1,
			want: "5-5-F-3-Z-2-E-9-W",
		},
		{
			s:    "5F3Z-2e-9-w",
			k:    4,
			want: "5F3Z-2E9W",
		},
		{
			s:    "2-5g-3-J",
			k:    2,
			want: "2-5G-3J",
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.s, tt.k,
		)
		t.Run(testName, func(t *testing.T) {
			got := licenseKeyFormatting(tt.s, tt.k)
			if got != tt.want {
				t.Errorf("licenseKeyFormatting(%v, %v) = %v, want %v", tt.s, tt.k, got, tt.want)
			}
		})
	}
}
