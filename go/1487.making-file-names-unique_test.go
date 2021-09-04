package main

import (
	"reflect"
	"testing"
)

func Test_getFolderNames(t *testing.T) {
	tests := []struct {
		name  string
		names []string
		want  []string
	}{
		{
			names: []string{"gta", "gta(1)", "gta", "avalon"},
			want:  []string{"gta", "gta(1)", "gta(2)", "avalon"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFolderNames(tt.names); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFolderNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
