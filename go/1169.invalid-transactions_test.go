package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_invalidTransactions(t *testing.T) {
	tests := []struct {
		name         string
		transactions []string
		want         []string
	}{
		{
			transactions: []string{"alex,676,260,bangkok", "bob,656,1366,bangkok", "alex,393,616,bangkok", "bob,820,990,amsterdam", "alex,596,1390,amsterdam"},
			want:         []string{"bob,656,1366,bangkok", "alex,596,1390,amsterdam"},
		},

		{
			transactions: []string{"bob,689,1910,barcelona", "alex,696,122,bangkok", "bob,832,1726,barcelona", "bob,820,596,bangkok", "chalicefy,217,669,barcelona", "bob,175,221,amsterdam"},
			want:         []string{"bob,689,1910,barcelona", "bob,832,1726,barcelona", "bob,820,596,bangkok"},
		},

		{
			transactions: []string{"alice,20,800,mtv", "bob,50,1200,mtv", "alice,20,800,mtv", "alice,50,1200,mtv", "alice,20,800,mtv", "alice,50,100,beijing"},
			want:         []string{"alice,20,800,mtv", "bob,50,1200,mtv", "alice,20,800,mtv", "alice,50,1200,mtv", "alice,20,800,mtv", "alice,50,100,beijing"},
		},

		{
			transactions: []string{"alice,20,800,mtv", "alice,50,100,mtv", "alice,51,100,frankfurt"},
			want:         []string{"alice,20,800,mtv", "alice,50,100,mtv", "alice,51,100,frankfurt"},
		},

		{
			transactions: []string{"alice,20,1220,mtv", "alice,20,1220,mtv"},
			want:         []string{"alice,20,1220,mtv", "alice,20,1220,mtv"},
		},

		{
			transactions: []string{"alice,20,800,mtv", "alice,50,100,beijing"},
			want:         []string{"alice,20,800,mtv", "alice,50,100,beijing"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := invalidTransactions(tt.transactions)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
