package main

import (
	"fmt"
	"testing"
)

func Test_minNonZeroProduct(t *testing.T) {
	tests := []struct {
		p    int
		want int
	}{
		{p: 32, want: 505517599},
		{p: 3, want: 1512},
		{p: 1, want: 1},
		{p: 2, want: 6},
		{p: 30, want: 945196305},
		{p: 31, want: 138191773},
		{p: 33, want: 861896614},
		{p: 34, want: 640964173},
		{p: 35, want: 112322054},
		{p: 36, want: 217659727},
		{p: 37, want: 680742062},
		{p: 38, want: 673217940},
		{p: 39, want: 945471045},
		{p: 40, want: 554966674},
		{p: 41, want: 190830260},
		{p: 42, want: 403329489},
		{p: 43, want: 305023508},
		{p: 44, want: 229675479},
		{p: 45, want: 865308368},
		{p: 46, want: 689473871},
		{p: 47, want: 161536946},
		{p: 48, want: 99452142},
		{p: 49, want: 720364340},
		{p: 50, want: 172386396},
		{p: 51, want: 198445540},
		{p: 52, want: 265347860},
		{p: 53, want: 504260931},
		{p: 54, want: 247773741},
		{p: 55, want: 65332879},
		{p: 56, want: 891336224},
		{p: 57, want: 221172799},
		{p: 58, want: 643213635},
		{p: 59, want: 926891661},
		{p: 60, want: 813987236},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.p), func(t *testing.T) {
			if got := minNonZeroProduct(tt.p); got != tt.want {
				t.Errorf("minNonZeroProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
