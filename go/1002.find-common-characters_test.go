package main

import (
	"reflect"
	"testing"
)

func Test_commonChars(t *testing.T) {
	tests := []struct {
		a    []string
		want []string
	}{
		{
			a: []string{"bella", "label", "roller"}, want: []string{"e", "l", "l"},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := commonChars(tt.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
