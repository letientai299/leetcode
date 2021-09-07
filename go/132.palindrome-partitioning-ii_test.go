package main

import "testing"

func Test_minCut(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "aaabaa", want: 1},
		{s: "kwtbjmsjvbrwriqwxadwnufplszhqccayvdhhvscxjaqsrmrrqngmuvxnugdzjfxeihogzsdjtvdmkudckjoggltcuybddbjoizu", want: 89},
		{
			s:    "adabdcaebdcebdcacaaaadbbcadabcbeabaadcbcaaddebdbddcbdacdbbaedbdaaecabdceddccbdeeddccdaabbabbdedaaabcdadbdabeacbeadbaddcbaacdbabcccbaceedbcccedbeecbccaecadccbdbdccbcbaacccbddcccbaedbacdbcaccdcaadcbaebebcceabbdcdeaabdbabadeaaaaedbdbcebcbddebccacacddebecabccbbdcbecbaeedcdacdcbdbebbacddddaabaedabbaaabaddcdaadcccdeebcabacdadbaacdccbeceddeebbbdbaaaaabaeecccaebdeabddacbedededebdebabdbcbdcbadbeeceecdcdbbdcbdbeeebcdcabdeeacabdeaedebbcaacdadaecbccbededceceabdcabdeabbcdecdedadcaebaababeedcaacdbdacbccdbcece",
			want: 273,
		},
		{s: "aaacbbbb", want: 2},
		{s: "aaabbbb", want: 1},
		{s: "bb", want: 0},
		{s: "aaab", want: 1},
		{s: "aaabaaa", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := minCut(tt.s); got != tt.want {
				t.Errorf("minCut() = %v, want %v", got, tt.want)
			}
		})
	}
}
