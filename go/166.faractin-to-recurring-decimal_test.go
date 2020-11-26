package main

import (
	"fmt"
	"testing"
)

func Test_fractionToDecimal(t *testing.T) {
	tests := []struct {
		n    int
		d    int
		want string
	}{
		{n: 1, d: 214748364, want: "0.00(000000465661289042462740251655654056577585848337359161441621040707904997124914069194026549138227660723878669455195477065427143370461252966751355553982241280310754777158628319049732085502639731402098131932683780538602845887105337854867197032523144157689601770377165713821223802198558308923834223016478952081795603341592860749337303449725)"},
		{n: 1, d: 9, want: "0.(1)"},
		{n: 7, d: -12, want: "-0.58(3)"},
		{n: 118, d: 7, want: "16.(857142)"},
		{n: 10, d: 3, want: "3.(3)"},
		{n: 10, d: 9, want: "1.(1)"},
		{n: 4, d: 333, want: "0.(012)"},
		{n: 4, d: 2, want: "2"},
		{n: -50, d: 8, want: "-6.25"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d/%d == %s", tt.n, tt.d, tt.want), func(t *testing.T) {
			t.Logf("%v", float64(tt.n)/float64(tt.d))
			if got := fractionToDecimal(tt.n, tt.d); got != tt.want {
				t.Errorf("fractionToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
