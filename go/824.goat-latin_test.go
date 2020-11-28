package main

import (
	"testing"
)

func Test_toGoatLatin(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			s:    "The quick brown fox jumped over the lazy dog",
			want: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		},
		{
			s:    "I speak Goat Latin",
			want: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
		},
		{s: "ab", want: "abmaa"},
		{s: "ba", want: "abmaa"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := toGoatLatin(tt.s); got != tt.want {
				t.Errorf("toGoatLatin() = %v, want %v", got, tt.want)
			}
		})
	}
}
