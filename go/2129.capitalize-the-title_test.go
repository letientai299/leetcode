package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_capitalizeTitle(t *testing.T) {
	tests := []struct {
		title string
		want  string
	}{
		{
			title: "L hV",
			want:  "l hv",
		},

		{
			title: "capiTalIze tHe titLe",
			want:  "Capitalize The Title",
		},

		{
			title: "First leTTeR of EACH Word",
			want:  "First Letter of Each Word",
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			assert.Equalf(t, tt.want, capitalizeTitle(tt.title), "capitalizeTitle(%v)", tt.title)
		})
	}
}
