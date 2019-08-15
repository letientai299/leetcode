package main

import (
	"fmt"
	"testing"
)

func Test_shortestCompletingWord(t *testing.T) {
	tests := []struct {
		licensePlate string
		words        []string
		want         string
	}{
		{
			licensePlate: "Ar16259",
			words:        []string{"our", "nature", "though", "party", "food", "any", "democratic", "building", "eat", "structure"},
			want:         "party",
		},

		{
			licensePlate: "1s3 PSt",
			words:        []string{"step", "steps", "stripe", "stepple"},
			want:         "steps",
		},

		{
			licensePlate: "1s3 456",
			words:        []string{"looks", "pest", "stew", "show"},
			want:         "pest",
		},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.licensePlate, tt.words,
		)
		t.Run(testName, func(t *testing.T) {
			got := shortestCompletingWord(tt.licensePlate, tt.words)
			if got != tt.want {
				t.Errorf("shortestCompletingWord(%v, %v) = %v, want %v", tt.licensePlate, tt.words, got, tt.want)
			}
		})
	}
}
