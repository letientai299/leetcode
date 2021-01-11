package main

import (
	"reflect"
	"testing"
)

func Test_findItinerary(t *testing.T) {
	tests := []struct {
		name    string
		tickets [][]string
		want    []string
	}{
		{
			tickets: [][]string{
				{"JFK", "ATL"}, {"ORD", "PHL"}, {"JFK", "ORD"}, {"PHX", "LAX"}, {"LAX", "JFK"}, {"PHL", "ATL"}, {"ATL", "PHX"},
			},
			want: []string{"JFK", "ATL", "PHX", "LAX", "JFK", "ORD", "PHL", "ATL"},
		},

		{
			tickets: [][]string{
				{"JFK", "AAA"}, {"AAA", "JFK"}, {"JFK", "BBB"}, {"JFK", "CCC"}, {"CCC", "JFK"},
			},
			want: []string{"JFK", "AAA", "JFK", "CCC", "JFK", "BBB"},
		},

		{
			tickets: [][]string{
				{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"},
			},
			want: []string{"JFK", "NRT", "JFK", "KUL"},
		},

		{
			tickets: [][]string{
				{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"},
			},
			want: []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},

		{
			tickets: [][]string{
				{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"},
			},
			want: []string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findItinerary(tt.tickets); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findItinerary() = %v, want %v", got, tt.want)
			}
		})
	}
}
