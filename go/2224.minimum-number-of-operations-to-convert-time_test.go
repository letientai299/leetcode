package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convertTime(t *testing.T) {
	tests := []struct {
		name    string
		current string
		correct string
		want    int
	}{
		{current: "09:41", correct: "10:34", want: 7},
		{current: "11:00", correct: "11:01", want: 1},
		{current: "02:30", correct: "04:35", want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, convertTime(tt.current, tt.correct), "convertTime(%v, %v)", tt.current, tt.correct)
		})
	}
}
