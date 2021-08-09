package main

import (
	"reflect"
	"testing"
)

func Test_alertNames(t *testing.T) {
	tests := []struct {
		name    string
		keyName []string
		keyTime []string
		want    []string
	}{
		{
			keyName: []string{"alice", "alice", "alice", "bob", "bob", "bob", "bob"},
			keyTime: []string{"12:01", "12:00", "18:00", "21:00", "21:20", "21:30", "23:00"},
			want:    []string{"bob"},
		},

		{
			keyName: []string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"},
			keyTime: []string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"},
			want:    []string{"daniel"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alertNames(tt.keyName, tt.keyTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("alertNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
