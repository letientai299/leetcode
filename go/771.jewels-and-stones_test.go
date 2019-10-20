package main

import (
	"fmt"
	"testing"
)

func Test_numJewelsInStones(t *testing.T) {
	tests := []struct {
		J    string
		S    string
		want int
	}{
		{J: "aA", S: "aAAbbbb", want: 3},
	}
	for _, tc := range tests {
		tt := tc
		testName := fmt.Sprintf(
			"%v, %v",
			tt.J, tt.S,
		)
		t.Run(testName, func(t *testing.T) {
			got := numJewelsInStones(tt.J, tt.S)
			if got != tt.want {
				t.Errorf("numJewelsInStones(%v, %v) = %v, want %v", tt.J, tt.S, got, tt.want)
			}
		})
	}
}
