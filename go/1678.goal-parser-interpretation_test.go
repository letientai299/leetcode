package main

import "testing"

func Test_interpret(t *testing.T) {
	tests := []struct {
		cmd  string
		want string
	}{
		{cmd: "G()(al)", want: "Goal"},
		{cmd: "()G()(al)", want: "oGoal"},
	}
	for _, tt := range tests {
		t.Run(tt.cmd, func(t *testing.T) {
			if got := interpret(tt.cmd); got != tt.want {
				t.Errorf("interpret() = %v, want %v", got, tt.want)
			}
		})
	}
}
