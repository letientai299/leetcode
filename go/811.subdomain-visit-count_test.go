package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subdomainVisits(t *testing.T) {
	tests := []struct {
		cpdomains []string
		want      []string
	}{
		{
			cpdomains: []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"},
			want:      []string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"},
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v",
			tt.cpdomains,
		)
		t.Run(testName, func(t *testing.T) {
			got := subdomainVisits(tt.cpdomains)
			assert.ElementsMatch(t, got, tt.want, "subdomainVisits(%v) = %v, want %v", tt.cpdomains, got, tt.want)
		})
	}
}
