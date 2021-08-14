package main

import (
	"reflect"
	"testing"
)

func Test_suggestedProducts(t *testing.T) {
	tests := []struct {
		name       string
		products   []string
		searchWord string
		want       [][]string
	}{
		{
			products: []string{
				"mobile", "mouse", "moneypot", "monitor", "mousepad",
			},

			searchWord: "mouse",
			want: [][]string{
				{"mobile", "moneypot", "monitor"},
				{"mobile", "moneypot", "monitor"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := suggestedProducts(tt.products, tt.searchWord); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("suggestedProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}
