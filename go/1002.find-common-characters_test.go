package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_commonChars(t *testing.T) {
	tests := []struct {
		a    []string
		want []string
	}{
		{
			a: []string{"bella", "label", "roller"}, want: []string{"e", "l", "l"},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := commonChars(tt.a); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("commonChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
