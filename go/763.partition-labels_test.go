package main

import (
	"reflect"
	"testing"
)

func Test_partitionLabels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int
	}{
		{s: "memeopeezeooeopzepezzmozzmoezozezzezopzeeomzmzmemopeoozpmpzzmpempommpmzzmeeooezmozomzzoezmeopzzozeompppmmvqxxddqqqvvqkkvdddkqqdkkkvxddvvvxqkvxxdqkkxxvdddkqdqqqxkvkdvvkvvqqvxxkddxdkkqxdkvdxxqqxxdkddxlflfrjcrllrrcljfrcrccjrfrfcrcrjcrlfffrrfcjrcfrfclfccccljjflrlrfjfjlflfrcrjllcjcjffrclrfrfcrfjlfcflrisbnynsttniinunitsussbibnnyybnntiybuuinutitinitinsiiisbtnbtsbinnuttnytuisbsnssbnnituibybtiynninisgawhhgawwhaaagggahhhwwawwwahawwhwhgaghggwhhahagagwgwhwaawaahwawhaggwgaawhgaaahaahwgwhwghgwaaaawhwghhgaghah", want: []int{105, 93, 99, 97, 106}},
		{s: "ntswuqqbidunnixxpoxxuuupotaatwdainsotwvpxpsdvdbwvbtdiptwtxnnbtqbdvnbowqitudutpsxsbbsvtipibqpvpnivottsxvoqqaqdxiviidivndvdtbvadnxboiqivpusuxaaqnqaobutdbpiosuitdnopoboivopaapadvqwwnnwvxndpxbapixaspwxxxvppoptqxitsvaaawxwaxtbxuixsoxoqdtopqqivaitnpvutzchkygjjgjkcfzjzrkmyerhgkglcyffezmehjcllmlrjghhfkfylkgyhyjfmljkzglkklykrjgrmzjyeyzrrkymccefggczrjflykclfhrjjckjlmglrmgfzlkkhffkjrkyfhegyykrzgjzcgjhkzzmzyejycfrkkekmhzjgggrmchkeclljlyhjkchmhjlehhejjyccyegzrcrerfzczfelzrlfylzleefgefgmzzlggmejjjygehmrczmkrc", want: []int{246, 254}},
		{s: "aeriabbbbzzzz", want: []int{5, 4, 4}},
		{s: "aeriab", want: []int{5, 1}},
		{s: "aeria", want: []int{5}},
		{s: "ababcbacadefegdehijhklij", want: []int{9, 7, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := partitionLabels(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
