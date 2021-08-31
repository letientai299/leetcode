package main

import (
	"reflect"
	"testing"
)

func Test_camelMatch(t *testing.T) {
	tests := []struct {
		queries []string
		pattern string
		want    []bool
	}{
		{
			pattern: "FB",
			queries: []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"},
			want:    []bool{true, false, true, true, false},
		},

		{
			queries: []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"},
			pattern: "FoBa",
			want:    []bool{true, false, true, false, false},
		},

		{
			queries: []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"},
			pattern: "FoBaT",
			want:    []bool{false, true, false, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.pattern, func(t *testing.T) {
			if got := camelMatch(tt.queries, tt.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("camelMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
