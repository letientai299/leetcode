package main

import "testing"

func Test_isLongPressedName(t *testing.T) {
	tests := []struct {
		name  string
		typed string
		want  bool
	}{
		{name: "saeed", typed: "ssaaedd", want: false},
		{name: "leel", typed: "leeel", want: true},
		{name: "alex", typed: "aallex", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLongPressedName(tt.name, tt.typed); got != tt.want {
				t.Errorf("isLongPressedName() = %v, want %v", got, tt.want)
			}
		})
	}
}
